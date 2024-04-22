package database
import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  )
  
  func database (){
    dsn := "host=localhost user=postgres password=secretkey dbname=postblog port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
  }
 