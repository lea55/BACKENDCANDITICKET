package paypal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
)

type paypal struct {
	config *Config
	http   *http.Client
}

func New() platform.Paypal {
	return &paypal{
		config: NewConfig(),
		http:   http.DefaultClient,
	}
}

func (p paypal) GeneratePaymentLink(data platform.RqGeneratePaypalLink) (platform.PypGeneratePaymentResp, error) {
	var result platform.PypGeneratePaymentResp

	value := fmt.Sprintf("%v", data.Value)

	rqData := RqSaveOrder{
		Intent: "CAPTURE",
		PurchaseUnits: []PurchaseUnit{
			{
				Description: data.Description,
				Amount: PurchaseUnitAmount{
					CurrencyCode: "USD",
					Value:        value,
				},
			},
		},
		ApplicationContext: ApplicationContext{
			BrandName:   "Canditickets",
			LandingPage: "LOGIN",
			UserAction:  "PAY_NOW",
			ReturnUrl:   p.config.RedirectUrl + "/" + data.Reference,
			CancelUrl:   "",
		},
	}

	jsonRqData, err := json.Marshal(&rqData)
	if err != nil {
		return result, err
	}

	url := p.config.ApiUrl + "/v2/checkout/orders"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRqData))
	if err != nil {
		return result, err
	}

	auth := p.getBasicAuth()
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.http.Do(req)
	if err != nil {
		return result, err
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return result, readErr
	}

	if resp.StatusCode != http.StatusCreated {
		return result, errors.New(string(body))
	}

	var responseData RsCreateOrder
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return result, err
	}

	var generatedLink string

	for _, v := range responseData.Links {
		if v.Rel == "approve" {
			generatedLink = v.Href
			break
		}
	}

	return platform.PypGeneratePaymentResp{
		ReferenceID:   responseData.ID,
		GeneratedLink: generatedLink,
	}, nil

}

func (p paypal) getBasicAuth() string {
	auth := p.config.ClientID + ":" + p.config.Secret
	basicStr := base64.StdEncoding.EncodeToString([]byte(auth))
	return basicStr
}
