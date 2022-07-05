package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lea55/BACKENDCANDITICKET/src/global/core"
	user "github.com/lea55/BACKENDCANDITICKET/src/packages/user/domain"
	"github.com/pkg/errors"
)

type AuthToken struct {
	config *core.AppConfig
}

func NewAuthToken(appConfig *core.AppConfig) *AuthToken {
	return &AuthToken{config: appConfig}
}

func (t AuthToken) GenerateUserToken(u user.Entity) (string, error) {
	payload := jwt.MapClaims{
		"email":    u.Email,
		"id":       u.ID,
		"rol_code": u.UserRol.Code,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString([]byte(t.config.SecretTokenKey))
	if err != nil {
		return "", errors.Wrap(err, "Error en la generación de token de autenticación")
	}
	return tokenStr, nil
}
