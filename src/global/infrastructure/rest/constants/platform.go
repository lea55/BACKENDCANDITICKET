package restconst

import (
	"github.com/lea55/BACKENDCANDITICKET/src/global/infrastructure/facebook"
	"github.com/lea55/BACKENDCANDITICKET/src/global/infrastructure/google"
	"github.com/lea55/BACKENDCANDITICKET/src/global/infrastructure/paypal"
	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
)

type AppPlatform struct {
	Facebook platform.Facebook
	Google   platform.Google
	Paypal   platform.Paypal
}

func GetAppPlatforms() *AppPlatform {
	return &AppPlatform{
		Facebook: facebook.NewFacebookPlat(),
		Google:   google.NewGooglePlat(),
		Paypal:   paypal.New(),
	}
}
