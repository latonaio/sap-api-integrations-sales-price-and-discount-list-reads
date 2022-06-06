package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-price-and-discount-list-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSalesPriceAndDiscountList(raw []byte, l *logger.Logger) ([]SalesPriceAndDiscountList, error) {
	pm := &responses.SalesPriceAndDiscountList{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPriceAndDiscountList. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	salesPriceAndDiscountList := make([]SalesPriceAndDiscountList, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesPriceAndDiscountList = append(salesPriceAndDiscountList, SalesPriceAndDiscountList{
			ObjectID:                      data.ObjectID,
			PriceDiscountListID:           data.PriceDiscountListID,
			Description:                   data.Description,
			DescriptionLanguageCode:       data.DescriptionLanguageCode,
			CurrencyCode:                  data.CurrencyCode,
			ValidityStartDate:             data.ValidityStartDate,
			ValidityEndDate:               data.ValidityEndDate,
			ReleaseStatusCode:             data.ReleaseStatusCode,
			IsBasePriceList:               data.IsBasePriceList,
			IsBasePriceListByProdCategory: data.IsBasePriceListByProdCategory,
			IsDistributionChainPriceList:  data.IsDistributionChainPriceList,
			IsDistributionChainPriceListByProdCategory: data.IsDistributionChainPriceListByProdCategory,
			IsAccountHierarchySpecificPriceList:        data.IsAccountHierarchySpecificPriceList,
			IsCustomerGroupSpecificPriceList:           data.IsCustomerGroupSpecificPriceList,
			IsCustomerSpecificPriceList:                data.IsCustomerSpecificPriceList,
			IsOverallCustomerDiscount:                  data.IsOverallCustomerDiscount,
			IsOverallCustomerGroupDiscount:             data.IsOverallCustomerGroupDiscount,
			IsAccountHierarchySpecificDiscountList:     data.IsAccountHierarchySpecificDiscountList,
			IsCustomerSpecificDiscountProducts:         data.IsCustomerSpecificDiscountProducts,
			IsCustomerSpecificDiscountProductCategory:  data.IsCustomerSpecificDiscountProductCategory,
			ProductCategoryID:                          data.ProductCategoryID,
			SalesOrgID:                                 data.SalesOrgID,
			DistributionChannel:                        data.DistributionChannel,
			CustomerGroup:                              data.CustomerGroup,
			AccountID:                                  data.AccountID,
			EntityLastChangedOn:                        data.EntityLastChangedOn,
			ETag:                                       data.ETag,
			ToInternalPriceDiscountListItems:           data.InternalPriceDiscountListItems.Deferred.URI,
		})
	}

	return salesPriceAndDiscountList, nil
}

func ConvertToInternalPriceDiscountListItems(raw []byte, l *logger.Logger) ([]InternalPriceDiscountListItems, error) {
	pm := &responses.ToInternalPriceDiscountListItems{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to InternalPriceDiscountListItems. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	internalPriceDiscountListItems := make([]InternalPriceDiscountListItems, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		internalPriceDiscountListItems = append(internalPriceDiscountListItems, InternalPriceDiscountListItems{
			ObjectID:            data.ObjectID,
			ParentObjectID:      data.ParentObjectID,
			PriceDiscountListID: data.PriceDiscountListID,
			Amount:              data.Amount,
			AmountCurrencyCode:  data.AmountCurrencyCode,
			PriceUnitContent:    data.PriceUnitContent,
			PriceUnitCode:       data.PriceUnitCode,
			Percent:             data.Percent,
			ProductID:           data.ProductID,
			AccountID:           data.AccountID,
			CustomerGroup:       data.CustomerGroup,
			ProductCategoryID:   data.ProductCategoryID,
			EntityLastChangedOn: data.EntityLastChangedOn,
			ETag:                data.ETag,
		})
	}

	return internalPriceDiscountListItems, nil
}
