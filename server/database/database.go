package database

import (
	"database/sql"
    _ "github.com/lib/pq"
	"log"
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
	
	db, err := sql.Open("postgres", "postgres://username:password@localhost/database_name?sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

}
