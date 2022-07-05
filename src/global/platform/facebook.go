package platform

type Facebook interface {
	ValidateUser(token string, ID string) error
}
