package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) headerBPGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		bpID := getHeaderBPGeneralExistenceConfKey(mapper, &header, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(bpID) {
				wg2.Done()
				return
			}
			res, err := c.bPGeneralExistenceConfRequest(bpID, mapper, input, existenceMap, mtx, log)
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
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) itemBPGeneralExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		bpID := getItemBPGeneralExistenceConfKey(mapper, &item, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(bpID) {
				wg2.Done()
				return
			}
			res, err := c.bPGeneralExistenceConfRequest(bpID, mapper, input, existenceMap, mtx, log)
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
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) bPGeneralExistenceConfRequest(bpID int, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
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
	req.BPGeneralReturn.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}

	return "", nil
}

func getHeaderBPGeneralExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) int {
	var bpID int

	switch mapper.Field {
	case "OwnerProductionPlantBusinessPartner":
		if header.OwnerProductionPlantBusinessPartner == nil {
			bpID = 0
		} else {
			bpID = *header.OwnerProductionPlantBusinessPartner
		}
	}

	return bpID
}

func getItemBPGeneralExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) int {
	var bpID int

	switch mapper.Field {
	case "ProductionPlantBusinessPartner":
		if item.ProductionPlantBusinessPartner == nil {
			bpID = 0
		} else {
			bpID = *item.ProductionPlantBusinessPartner
		}
	}
	return bpID
}
