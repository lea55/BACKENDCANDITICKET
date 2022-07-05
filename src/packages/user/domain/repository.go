package user

type Repository interface {
	Save(user Entity) error
	FindByEmail(email string) (Entity, error)
	FindEnabled() ([]Entity, error)
	FindByID(ID string) (Entity, error)
	UpdateImage(userID string, imageUrl string) error
	UpdateEmail(ID string, email string) error
	UpdatePassword(ID string, pass string) error
	UpdateDeviceID(userID string, deviceID string) error
}
