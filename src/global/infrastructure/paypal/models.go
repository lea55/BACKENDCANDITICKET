package paypal

type RqSaveOrder struct {
	Intent             string             `json:"intent"`
	PurchaseUnits      []PurchaseUnit     `json:"purchase_units"`
	ApplicationContext ApplicationContext `json:"application_context"`
}
type PurchaseUnit struct {
	Description string             `json:"description"`
	Amount      PurchaseUnitAmount `json:"amount"`
}

type PurchaseUnitAmount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type ApplicationContext struct {
	BrandName   string `json:"brand_name"`
	LandingPage string `json:"landing_page"`
	UserAction  string `json:"user_action"`
	ReturnUrl   string `json:"return_url"`
	CancelUrl   string `json:"cancel_url"`
}

type RsCreateOrder struct {
	ID     string              `json:"id"`
	Status string              `json:"status"`
	Links  []RsCreateOrderLink `json:"links"`
}

type RsCreateOrderLink struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
