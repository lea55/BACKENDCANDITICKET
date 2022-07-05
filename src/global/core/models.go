package core

import (
	"os"
)

var appConfig *AppConfig

type AppConfig struct {
	BaseUrl            string
	SecretTokenKey     string
	AssetsImageDefault string
	DocsPath           string
	ImagePath          string
	ServerUrl          string
	MigrationEnabled   bool
	MongoCnn           string
	MongoDBName        string
}

func GetAppConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	} else {
		setAppConfig()
		return appConfig
	}

}

func setAppConfig() {
	appConfig = &AppConfig{
		BaseUrl:            os.Getenv("BASE_URL"),
		SecretTokenKey:     os.Getenv("SECRET_TOKEN_KEY"),
		AssetsImageDefault: os.Getenv("ASSETS_IMAGE_DEFAULT"),
		DocsPath:           os.Getenv("DOCS_PATH"),
		ImagePath:          os.Getenv("IMAGE_PATH"),
		ServerUrl:          os.Getenv("SERVER_URL"),
		MigrationEnabled:   len(os.Getenv("MIGRATION_ENABLED")) > 0,
		MongoCnn:           os.Getenv("MONGO_CNN"),
		MongoDBName:        os.Getenv("MONGO_DB_NAME"),
	}
}
