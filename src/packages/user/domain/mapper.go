package user

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m Mapper) ConvertToBasic(user Entity) BasicUser {
	return BasicUser{
		ID:          user.ID,
		Names:       user.Name,
		Surname:     user.Surname,
		Code:        user.Code,
		Email:       user.Email,
		PhoneNumber: "",
		RolCode:     user.UserRol.Code,
	}
}
