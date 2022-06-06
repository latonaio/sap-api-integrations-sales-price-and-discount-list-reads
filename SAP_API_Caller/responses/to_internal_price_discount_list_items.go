package responses

type ToInternalPriceDiscountListItems struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
