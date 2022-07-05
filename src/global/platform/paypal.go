package platform

type RqGeneratePaypalLink struct {
	Value        float32
	CurrencyCode string
	Description  string
	Quantity     uint32
	Reference    string
}

type PypGeneratePaymentResp struct {
	ReferenceID   string
	GeneratedLink string
}

type Paypal interface {
	GeneratePaymentLink(data RqGeneratePaypalLink) (PypGeneratePaymentResp, error)
}
