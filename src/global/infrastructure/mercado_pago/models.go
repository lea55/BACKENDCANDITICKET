package mercpago

import "encoding/json"

func unmarshalPreferenceCreated(data []byte) (preferenceCreated, error) {
	var r preferenceCreated
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *preferenceCreated) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type preferenceCreated struct {
	AdditionalInfo     string                   `json:"additional_info"`
	AutoReturn         string                   `json:"auto_return"`
	BackUrls           preferenceUrls           `json:"back_urls"`
	BinaryMode         bool                     `json:"binary_mode"`
	ClientID           string                   `json:"client_id"`
	CollectorID        int64                    `json:"collector_id"`
	CouponCode         interface{}              `json:"coupon_code"`
	CouponLabels       interface{}              `json:"coupon_labels"`
	DateCreated        string                   `json:"date_created"`
	DateOfExpiration   interface{}              `json:"date_of_expiration"`
	ExpirationDateFrom interface{}              `json:"expiration_date_from"`
	ExpirationDateTo   interface{}              `json:"expiration_date_to"`
	Expires            bool                     `json:"expires"`
	ExternalReference  string                   `json:"external_reference"`
	ID                 string                   `json:"id"`
	InitPoint          string                   `json:"init_point"`
	InternalMetadata   interface{}              `json:"internal_metadata"`
	Items              []preferenceItem         `json:"items"`
	Marketplace        string                   `json:"marketplace"`
	MarketplaceFee     int64                    `json:"marketplace_fee"`
	Metadata           preferenceMetadata       `json:"metadata"`
	NotificationURL    string                   `json:"notification_url"`
	OperationType      string                   `json:"operation_type"`
	Payer              preferencePayer          `json:"payer"`
	PaymentMethods     preferencePaymentMethods `json:"payment_methods"`
	ProcessingModes    interface{}              `json:"processing_modes"`
	ProductID          interface{}              `json:"product_id"`
	RedirectUrls       preferenceUrls           `json:"redirect_urls"`
	SandboxInitPoint   string                   `json:"sandbox_init_point"`
	SiteID             string                   `json:"site_id"`
	Shipments          preferenceShipments      `json:"shipments"`
	TotalAmount        interface{}              `json:"total_amount"`
	LastUpdated        interface{}              `json:"last_updated"`
}

type preferenceUrls struct {
	Failure string `json:"failure"`
	Pending string `json:"pending"`
	Success string `json:"success"`
}

type preferenceItem struct {
	ID          string `json:"id"`
	CategoryID  string `json:"category_id"`
	CurrencyID  string `json:"currency_id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Quantity    int64  `json:"quantity"`
	UnitPrice   int64  `json:"unit_price"`
}

type preferenceMetadata struct {
}

type preferencePayer struct {
	Phone          preferencePhone          `json:"phone"`
	Address        preferenceAddress        `json:"address"`
	Email          string                   `json:"email"`
	Identification preferenceIdentification `json:"identification"`
	Name           string                   `json:"name"`
	Surname        string                   `json:"surname"`
	DateCreated    interface{}              `json:"date_created"`
	LastPurchase   interface{}              `json:"last_purchase"`
}

type preferenceAddress struct {
	ZipCode      string      `json:"zip_code"`
	StreetName   string      `json:"street_name"`
	StreetNumber interface{} `json:"street_number"`
}

type preferenceIdentification struct {
	Number string `json:"number"`
	Type   string `json:"type"`
}

type preferencePhone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type preferencePaymentMethods struct {
	DefaultCardID          interface{}                 `json:"default_card_id"`
	DefaultPaymentMethodID interface{}                 `json:"default_payment_method_id"`
	ExcludedPaymentMethods []preferenceExcludedPayment `json:"excluded_payment_methods"`
	ExcludedPaymentTypes   []preferenceExcludedPayment `json:"excluded_payment_types"`
	Installments           interface{}                 `json:"installments"`
	DefaultInstallments    interface{}                 `json:"default_installments"`
}

type preferenceExcludedPayment struct {
	ID string `json:"id"`
}

type preferenceShipments struct {
	DefaultShippingMethod interface{}               `json:"default_shipping_method"`
	ReceiverAddress       preferenceReceiverAddress `json:"receiver_address"`
}

type preferenceReceiverAddress struct {
	ZipCode      string      `json:"zip_code"`
	StreetName   string      `json:"street_name"`
	StreetNumber interface{} `json:"street_number"`
	Floor        string      `json:"floor"`
	Apartment    string      `json:"apartment"`
	CityName     interface{} `json:"city_name"`
	StateName    interface{} `json:"state_name"`
	CountryName  interface{} `json:"country_name"`
}

func unmarshalCreatePreferenceModel(data []byte) (createPreferenceModel, error) {
	var r createPreferenceModel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *createPreferenceModel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type createPreferenceModel struct {
	NotificationURL string                      `json:"notification_url"`
	Items           []createPreferenceItemModel `json:"items"`
}

type createPreferenceItemModel struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quantity    uint8   `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
}
