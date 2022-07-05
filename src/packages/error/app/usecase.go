package apperrlog

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UseCase struct {
	repo errlog.Repository
}

func NewUseCase(repo errlog.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc UseCase) Save(model RqSaveErr) error {

	newErr := errlog.Entity{
		ID:          uuid.NewString(),
		Description: model.Description,
		Module:      model.Module,
		Date:        time.Now().UTC(),
		UserID:      model.UserID,
	}

	err := uc.repo.Save(newErr)
	if err != nil {
		return errors.Wrap(err, "Error al guardar LOG de error")
	}

	return nil

}
