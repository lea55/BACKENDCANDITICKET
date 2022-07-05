package user

import (
	"time"

	rol "github.com/lea55/BACKENDCANDITICKET/src/packages/rol/domain"
)

type Entity struct {
	ID           string     `json:"uu_id"`
	Code         string     `json:"code"`
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	UserName     string     `json:"user_name"`
	UserRol      rol.Entity `json:"user_rol"`
	Enabled      bool       `json:"enabled"`
	PhoneNumber  string     `json:"phone_number"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ProfileImage string     `json:"profile_image"`
	DeviceID     string     `json:"device_id"`
}

type BasicUser struct {
	ID           string `json:"id,omitempty"`
	Names        string `json:"names,omitempty"`
	Surname      string `json:"surname,omitempty"`
	Code         string `json:"code,omitempty"`
	Email        string `json:"email,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	RolCode      string `json:"rol_code,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"`
}

func (usr Entity) GetBasic() BasicUser {
	return BasicUser{
		ID:           usr.ID,
		Names:        usr.Name,
		Surname:      usr.Surname,
		Code:         usr.Code,
		Email:        usr.Password,
		PhoneNumber:  usr.PhoneNumber,
		RolCode:      usr.UserRol.Code,
		ProfileImage: usr.ProfileImage,
	}
}
