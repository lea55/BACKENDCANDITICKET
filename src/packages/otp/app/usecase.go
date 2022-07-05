package appotp

import (
	"time"

	"github.com/google/uuid"
	"github.com/lea55/BACKENDCANDITICKET/src/global/helpers"
	otp "github.com/lea55/BACKENDCANDITICKET/src/packages/otp/domain"
	"github.com/pkg/errors"
)

type UseCase struct {
	repo     otp.Repository
	codeHelp *helpers.Code
}

func NewUseCase(repo otp.Repository) *UseCase {
	return &UseCase{repo: repo, codeHelp: helpers.NewCode()}
}

func (uc UseCase) Validate(code string) error {
	founded, err := uc.repo.FindByCode(code)
	if err != nil {
		return errors.Wrap(err, "El código enviado no es válido")
	}

	if !founded.Enabled {
		return errors.Wrap(err, "El vódigo ya no es válido")
	}

	now := time.Now().UTC()
	if now.After(founded.ExpireAt) {
		return errors.New("El código ya ha vencido, intentelo de nuevo")
	}

	return nil
}

func (uc UseCase) Generate(data RqGenerate) (string, error) {
	ID := uuid.NewString()
	code := uc.codeHelp.Random("", 100000, 999999)

	newItem := otp.Entity{
		ID:        ID,
		Code:      code,
		Type:      data.Type,
		CreatedAt: time.Now().UTC(),
		ExpireAt:  time.Now().UTC().Add(time.Minute * time.Duration(5)),
		Enabled:   true,
		UserID:    data.UserID,
	}

	err := uc.repo.Save(newItem)
	if err != nil {
		return "", errors.Wrap(err, "Error en la generación de código")
	}

	return code, nil
}
