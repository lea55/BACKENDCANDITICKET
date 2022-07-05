package otp

type Repository interface {
	Save(data Entity) error
	FindByCode(code string) (Entity, error)
}
