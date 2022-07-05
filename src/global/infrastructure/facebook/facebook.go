package facebook

import (
	"encoding/json"
	"net/http"

	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
	"github.com/pkg/errors"
)

type facebookPlat struct{}

type fbValidUser struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func NewFacebookPlat() platform.Facebook {
	return &facebookPlat{}
}

func (f facebookPlat) ValidateUser(token string, ID string) error {
	var validUser fbValidUser

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + token)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New("Token de facebok no valido")
	}

	err = json.NewDecoder(resp.Body).Decode(&validUser)
	if err != nil {
		return errors.Wrap(err, "Error en validacion de datos")
	}

	if ID != validUser.ID {
		return errors.New("Id de usuario no valido")
	}

	return nil
}
