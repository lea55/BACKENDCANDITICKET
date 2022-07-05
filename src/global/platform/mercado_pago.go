package platform

type MpGeneratePayment struct {
	Title              string
	Description        string
	Quantity           uint8
	UnitPrice          float32
	RedirectParamValue string
}

type MpGeneratePaymentResp struct {
	ReferenceID   string
	GeneratedLink string
}

type MercadoPago interface {
	GenerateProductPayment(model MpGeneratePayment) (MpGeneratePaymentResp, error)
}
