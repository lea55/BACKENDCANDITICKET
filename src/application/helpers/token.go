package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/lea55/BACKENDCANDITICKET/src/domain/config"
	"github.com/lea55/BACKENDCANDITICKET/src/domain/entity"
	"github.com/pkg/errors"
	"time"
)

type AuthToken struct {
	config *config.AppConfig
}

func NewAuthToken(appConfig *config.AppConfig) *AuthToken {
	return &AuthToken{config: appConfig}
}

func (t AuthToken) GenerateUserToken(user entity.User) (string, error) {
	payload := jwt.MapClaims{
		"email":    user.Email,
		"id":       user.ID,
		"rol_code": user.Rol.Code,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString([]byte(t.config.SecretTokenKey))
	if err != nil {
		return "", errors.Wrap(err, "Error en la generación de token de autenticación")
	}
	return tokenStr, nil
}
