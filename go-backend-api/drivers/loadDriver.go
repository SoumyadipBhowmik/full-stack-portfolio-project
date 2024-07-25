package drivers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var host = os.Getenv("DB_HOST")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var dbname = os.Getenv("DB_NAME")
var port = os.Getenv("DB_PORT")
var sslmode = os.Getenv("SSLMODE")

func ConnectToDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("unable to connect to database")
	}
}
