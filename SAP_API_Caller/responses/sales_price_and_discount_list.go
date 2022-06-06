package responses

type SalesPriceAndDiscountList struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                   string `json:"ObjectID"`
			PriceDiscountListID                        string `json:"PriceDiscountListID"`
			Description                                string `json:"Description"`
			DescriptionLanguageCode                    string `json:"DescriptionLanguageCode"`
			CurrencyCode                               string `json:"CurrencyCode"`
			ValidityStartDate                          string `json:"ValidityStartDate"`
			ValidityEndDate                            string `json:"ValidityEndDate"`
			ReleaseStatusCode                          string `json:"ReleaseStatusCode"`
			IsBasePriceList                            bool   `json:"IsBasePriceList"`
			IsBasePriceListByProdCategory              bool   `json:"IsBasePriceListByProdCategory"`
			IsDistributionChainPriceList               bool   `json:"IsDistributionChainPriceList"`
			IsDistributionChainPriceListByProdCategory bool   `json:"IsDistributionChainPriceListByProdCategory"`
			IsAccountHierarchySpecificPriceList        bool   `json:"IsAccountHierarchySpecificPriceList"`
			IsCustomerGroupSpecificPriceList           bool   `json:"IsCustomerGroupSpecificPriceList"`
			IsCustomerSpecificPriceList                bool   `json:"IsCustomerSpecificPriceList"`
			IsOverallCustomerDiscount                  bool   `json:"IsOverallCustomerDiscount"`
			IsOverallCustomerGroupDiscount             bool   `json:"IsOverallCustomerGroupDiscount"`
			IsAccountHierarchySpecificDiscountList     bool   `json:"IsAccountHierarchySpecificDiscountList"`
			IsCustomerSpecificDiscountProducts         bool   `json:"IsCustomerSpecificDiscountProducts"`
			IsCustomerSpecificDiscountProductCategory  bool   `json:"IsCustomerSpecificDiscountProductCategory"`
			ProductCategoryID                          string `json:"ProductCategoryID"`
			SalesOrgID                                 string `json:"SalesOrgID"`
			DistributionChannel                        string `json:"DistributionChannel"`
			CustomerGroup                              string `json:"CustomerGroup"`
			AccountID                                  string `json:"AccountID"`
			EntityLastChangedOn                        string `json:"EntityLastChangedOn"`
			ETag                                       string `json:"ETag"`
			InternalPriceDiscountListItems             struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"InternalPriceDiscountListItems"`
		} `json:"results"`
	} `json:"d"`
}
