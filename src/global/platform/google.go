package platform

type Google interface {
	ValidateUser(token string, ID string) error
}
