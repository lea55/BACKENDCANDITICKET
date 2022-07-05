package mercpago

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
)

type mercadoPagoImpl struct {
	AccessToken     string
	PublicKey       string
	Url             string
	NotificationUrl string
	httpClient      *http.Client
	env             string
}

func NewMercadoPagoImpl() platform.MercadoPago {
	return &mercadoPagoImpl{
		AccessToken:     os.Getenv("MERCADO_PAGO_ACCESS_TOKEN"),
		PublicKey:       os.Getenv("MERCADO_PAGO_PUBLIC_KEY"),
		Url:             os.Getenv("MERCADO_PAGO_BASE_URL"),
		NotificationUrl: os.Getenv("MERCADO_PAGO_NOTIFICATION_BASE_URL"),
		env:             os.Getenv("ENV"),
		httpClient:      http.DefaultClient,
	}
}

func (m mercadoPagoImpl) GenerateProductPayment(model platform.MpGeneratePayment) (platform.MpGeneratePaymentResp, error) {
	var responseData platform.MpGeneratePaymentResp
	url := m.Url + "checkout/preferences?access_token=" + m.AccessToken

	items := make([]createPreferenceItemModel, 0)

	newItem := createPreferenceItemModel{
		Title:       model.Title,
		Description: model.Description,
		Quantity:    model.Quantity,
		UnitPrice:   model.UnitPrice,
	}

	items = append(items, newItem)

	data := createPreferenceModel{
		NotificationURL: m.NotificationUrl + model.RedirectParamValue,
		Items:           items,
	}

	jsonRqData, err := data.Marshal()
	if err != nil {
		return responseData, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRqData))
	if err != nil {
		return responseData, err
	}

	req.Header.Add("Authorization", "Bearer"+" "+m.PublicKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return responseData, err
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return responseData, err
	}

	if resp.StatusCode != http.StatusCreated {
		return responseData, errors.New(string(body))
	}

	response, err := unmarshalPreferenceCreated(body)

	var initPoint string

	if m.env == "PROD" {
		initPoint = response.InitPoint
	} else {
		initPoint = response.SandboxInitPoint
	}

	responseData = platform.MpGeneratePaymentResp{
		ReferenceID:   response.ID,
		GeneratedLink: initPoint,
	}

	return responseData, nil

}
