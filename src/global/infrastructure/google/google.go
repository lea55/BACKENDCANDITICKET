package google

import (
	"encoding/json"
	"net/http"

	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
	"github.com/pkg/errors"
)

type googlePlat struct{}

type googleValidUser struct {
	IssuedTo      string `json:"issued_to"`
	Audience      string `json:"audience"`
	UserID        string `json:"user_id"`
	Scope         string `json:"scope"`
	ExpiresIn     int64  `json:"expires_in"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	AccessType    string `json:"access_type"`
}

func NewGooglePlat() platform.Google {
	return &googlePlat{}
}

func (g googlePlat) ValidateUser(token string, ID string) error {
	var validUser googleValidUser

	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=" + token)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New("Token de google no valido")
	}

	err = json.NewDecoder(resp.Body).Decode(&validUser)
	if err != nil {
		return errors.Wrap(err, "Error en validacion de datos")
	}

	if ID != validUser.UserID {
		return errors.New("Id de usuario no valido")
	}

	return nil
}
