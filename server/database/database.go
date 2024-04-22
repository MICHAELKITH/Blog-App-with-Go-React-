package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)
 
var DBconn *gorm.DB
  func ConnectDB (){
    dsn := "host=localhost user=postgres password=secretkey dbname=postblog port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    defer db.close()
    if err != nil {
      log.Fatal(err)
    }

    if err = db.Ping(); err != nil{
      log.Fatal(err)
    }
    
  }
 