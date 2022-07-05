package rol

import "github.com/lea55/BACKENDCANDITICKET/src/global/core"

type Repository interface {
	FindAll() ([]Entity, error)
	FindByCode(code core.UserRol) (Entity, error)
}
