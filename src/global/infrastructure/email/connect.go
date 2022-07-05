package email

import (
	"log"
	"strconv"

	mail "github.com/xhit/go-simple-mail"
)

var ec *mail.SMTPClient

func GetClient() *mail.SMTPClient {
	config := NewConfigFromEnv()

	port, err := strconv.Atoi(config.Port)
	if err != nil {
		log.Fatal("Puerto host email no válido")
	}

	server := mail.NewSMTPClient()
	server.Host = config.Host
	server.Port = port
	server.Username = config.UserName
	server.Password = config.Password
	server.Encryption = mail.EncryptionTLS

	result, err := server.Connect()
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Servidor email conectado correctamente")
	ec = result
	return ec
}
