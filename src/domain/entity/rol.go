package entity

type Rol struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RolCode struct {
	Organizer string
	User      string
	Support   string
}

func NewUserRol() *RolCode {
	return &RolCode{Organizer: "ORGANIZER", User: "USER", Support: "SUPPORT"}
}
