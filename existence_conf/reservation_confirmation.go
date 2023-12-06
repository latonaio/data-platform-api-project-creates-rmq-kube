package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) itemComponentReservationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		for _, itemComponent := range item.ItemComponent {
			reservation := getItemComponentReservationExistenceConfKey(mapper, &itemComponent, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(reservation) {
					wg2.Done()
					return
				}
				res, err := c.reservationExistenceConfRequest(reservation, mapper, input, existenceMap, mtx, log)
				if err != nil {
					mtx.Lock()
					*errs = append(*errs, err)
					mtx.Unlock()
				}
				if res != "" {
					*exconfErrMsg = res
				}
				wg2.Done()
			}()
		}
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) reservationExistenceConfRequest(reservation int, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"Reservation": reservation,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.ReservationReturn.Reservation = reservation

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}

	return "", nil
}

func getItemComponentReservationExistenceConfKey(mapper ExConfMapper, itemComponent *dpfm_api_input_reader.ItemComponent, exconfErrMsg *string) int {
	var reservation int

	switch mapper.Field {
	case "reservation":
		if itemComponent.Reservation == nil {
			reservation = 0
		} else {
			reservation = *itemComponent.Reservation
		}
	}
	return reservation
}
