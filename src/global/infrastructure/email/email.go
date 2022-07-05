package email

import (
	"fmt"

	"github.com/lea55/BACKENDCANDITICKET/src/global/platform"
	mail "github.com/xhit/go-simple-mail"
)

type emailImpl struct {
	config *Config
	client *mail.SMTPClient
}

func NewEmailImpl() platform.Email {
	return &emailImpl{
		config: NewConfigFromEnv(),
		client: GetClient(),
	}
}

func (e emailImpl) SendEmail(data platform.RqSendEmail) error {

	email := mail.NewMSG()
	email.SetFrom("Canditickets <" + e.config.UserName + ">")
	email.AddTo(data.EmailDestination)
	email.SetSubject(data.Subject)

	body := e.buildHtmlEmail(data.Tittle, data.Content)

	email.SetBody(mail.TextHTML, body)

	client := GetClient()

	err := email.Send(client)
	if err != nil {
		return err
	}

	return nil
}

func (e emailImpl) buildHtmlEmail(title string, content []string) string {

	var body string

	for _, v := range content {
		item := "<p> " + v + " </p>      "
		body = body + item
	}

	result := fmt.Sprintf(`
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
<h1> %s </h1>
<br>
  %s
</body>`, title, body)

	return result
}
