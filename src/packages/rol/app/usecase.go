package roluc

import (
	rol "github.com/lea55/BACKENDCANDITICKET/src/packages/rol/domain"
	"github.com/pkg/errors"
)

type UseCase struct {
	repo rol.Repository
}

func NewUseCase(repo rol.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc UseCase) GetAll() ([]rol.Entity, error) {
	roles, err := uc.repo.FindAll()
	if err != nil {
		return roles, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	return roles, nil
}
