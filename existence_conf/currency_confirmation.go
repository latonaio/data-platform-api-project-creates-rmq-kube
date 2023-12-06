package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) itemComponentCostingCurrencyExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		for _, itemComponent := range item.ItemComponent {
			for _, itemComponentCosting := range itemComponent.ItemComponentCosting {
				currency := getItemComponentCostingCurrencyExistenceConfKey(mapper, &itemComponentCosting, exconfErrMsg)
				wg2.Add(1)
				exReqTimes++
				go func() {
					if isZero(currency) {
						wg2.Done()
						return
					}
					res, err := c.currencyExistenceConfRequest(currency, mapper, input, existenceMap, mtx, log)
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
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) currencyExistenceConfRequest(currency string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"Currency": currency,
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
	req.CurrencyReturn.Currency = currency

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getItemComponentCostingCurrencyExistenceConfKey(mapper ExConfMapper, itemComponentCosting *dpfm_api_input_reader.ItemComponentCosting, exconfErrMsg *string) string {
	var currency string

	switch mapper.Field {
	case "Currency":
		if itemComponentCosting.Currency == nil {
			currency = ""
		} else {
			currency = *itemComponentCosting.Currency
		}
	}
	return currency
}
