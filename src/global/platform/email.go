package platform

type RqSendEmail struct {
	EmailDestination string   `binding:"required,email"`
	Subject          string   `binding:"required"`
	Tittle           string   `binding:"required"`
	Content          []string `binding:"required"`
}

type Email interface {
	SendEmail(data RqSendEmail) error
}
