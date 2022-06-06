package main

import (
	sap_api_caller "sap-api-integrations-sales-price-and-discount-list-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-sales-price-and-discount-list-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Sales_Price_And_Discount_list_Sales_Price_And_Discount_list_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/sap/c4c/odata/v1/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"SalesPriceAndDiscountList",
		}
	}

	caller.AsyncGetSalesPriceAndDiscountList(
		inoutSDC.SalesPriceAndDiscountList.PriceDiscountListID,
		accepter,
	)
}
