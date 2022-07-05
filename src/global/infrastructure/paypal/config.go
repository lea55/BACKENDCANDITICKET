package paypal

import "os"

type Config struct {
	ClientID    string
	Account     string
	Secret      string
	ApiUrl      string
	RedirectUrl string
}

func NewConfig() *Config {
	return &Config{
		RedirectUrl: os.Getenv("PAYPAL_REDIRECT_URL"),
		ClientID:    os.Getenv("PAYPAL_CLIENT_ID"),
		Account:     os.Getenv("PAYPAL_CLIENT_ACCOUNT"),
		Secret:      os.Getenv("PAYPAL_CLIENT_SECRET"),
		ApiUrl:      os.Getenv("PAYPAL_API"),
	}
}
