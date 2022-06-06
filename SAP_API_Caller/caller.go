package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-sales-price-and-discount-list-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesPriceAndDiscountList(priceDiscountListID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "SalesPriceAndDiscountList":
			func() {
				c.SalesPriceAndDiscountList(priceDiscountListID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) SalesPriceAndDiscountList(priceDiscountListID string) {
	salesPriceAndDiscountListData, err := c.callSalesPriceAndDiscountListSrvAPIRequirementSalesPriceAndDiscountList("InternalPriceDiscountListCollection", priceDiscountListID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(salesPriceAndDiscountListData)

	internalPriceDiscountListItemsData, err := c.callInternalPriceDiscountListItems(salesPriceAndDiscountListData[0].ToInternalPriceDiscountListItems)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(internalPriceDiscountListItemsData)

}

func (c *SAPAPICaller) callSalesPriceAndDiscountListSrvAPIRequirementSalesPriceAndDiscountList(api, priceDiscountListID string) ([]sap_api_output_formatter.SalesPriceAndDiscountList, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithSalesPriceAndDiscountList(req, priceDiscountListID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesPriceAndDiscountList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callInternalPriceDiscountListItems(url string) ([]sap_api_output_formatter.InternalPriceDiscountListItems, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToInternalPriceDiscountListItems(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithSalesPriceAndDiscountList(req *http.Request, priceDiscountListID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PriceDiscountListID eq '%s'", priceDiscountListID))
	req.URL.RawQuery = params.Encode()
}
