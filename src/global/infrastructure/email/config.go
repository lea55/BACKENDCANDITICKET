package email

import "os"

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
}

func NewConfigFromEnv() *Config {
	return &Config{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		UserName: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}
}
