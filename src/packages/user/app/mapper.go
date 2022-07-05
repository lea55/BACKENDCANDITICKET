package useruc

import (
	"time"

	"github.com/lea55/BACKENDCANDITICKET/packages/rol/domain"
	rol "github.com/lea55/BACKENDCANDITICKET/src/packages/rol/domain"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m Mapper) userToRegister(
	rol *rol.Entity,
	pass string,
	userName string,
	uuId string,
	code string,
	dto *RqRegisterUser,
) user.Entity {
	return user.Entity{
		ID:        uuId,
		Code:      code,
		Name:      dto.Name,
		Surname:   dto.Surname,
		Email:     dto.Email,
		Password:  pass,
		UserName:  userName,
		UserRol:   *rol,
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m Mapper) BasicUser(u *user.Entity) user.BasicUser {
	return user.BasicUser{
		ID:          u.ID,
		Names:       u.Name,
		Surname:     u.Surname,
		Code:        u.Code,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		RolCode:     u.UserRol.Code,
	}
}
