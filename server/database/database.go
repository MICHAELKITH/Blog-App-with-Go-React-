package database

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBname   string
	SSLMode  string
}

var DBconn *gorm.DB

func ConnectDB() {
	dsn := " user=postgres password=secretkey dbname=postblog port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	defer db.close()
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}
