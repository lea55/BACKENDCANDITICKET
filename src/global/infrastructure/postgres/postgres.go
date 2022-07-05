package apppsql

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	db = connect()
}

func GetCnn() *gorm.DB {
	if db == nil {
		db = connect()
	}

	return db
}

func connect() *gorm.DB {
	host := os.Getenv("PSQL_HOST")
	user := os.Getenv("PSQL_USER")
	pass := os.Getenv("PSQL_PASSWORD")
	dbName := os.Getenv("PSQL_DEFAULT_DB")
	port := os.Getenv("PSQL_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, pass, dbName, port)
	dbCnn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return dbCnn
}
