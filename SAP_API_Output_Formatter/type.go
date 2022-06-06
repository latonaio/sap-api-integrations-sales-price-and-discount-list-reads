package sap_api_output_formatter

type SalesPriceAndDiscountListReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	DiscountCode  string `json:"discount_code"`
	Deleted       bool   `json:"deleted"`
}

type SalesPriceAndDiscountList struct {
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
	ToInternalPriceDiscountListItems           string `json:"InternalPriceDiscountListItems"`
}
type InternalPriceDiscountListItems struct {
	ObjectID            string `json:"ObjectID"`
	ParentObjectID      string `json:"ParentObjectID"`
	PriceDiscountListID string `json:"PriceDiscountListID"`
	Amount              string `json:"Amount"`
	AmountCurrencyCode  string `json:"AmountCurrencyCode"`
	PriceUnitContent    string `json:"PriceUnitContent"`
	PriceUnitCode       string `json:"PriceUnitCode"`
	Percent             string `json:"Percent"`
	ProductID           string `json:"ProductID"`
	AccountID           string `json:"AccountID"`
	CustomerGroup       string `json:"CustomerGroup"`
	ProductCategoryID   string `json:"ProductCategoryID"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
	ETag                string `json:"ETag"`
}
