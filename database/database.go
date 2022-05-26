package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err errors



dialect := os.Getenv("DIALECT")
host := os.Getenv("HOST")
dbPort := os.Getenv("DBPORT")
user := os.Getenv("USER")
dbName := os.Getenv("NAME")
dbpassword := os.Getenv("PASSWORD")



dbURI := fmt.Sprintf("host=%s user=%s dbname%s sslmode=disable password=%s port=%s", host, user, dbName, db
password, dbPort)



db, err = gorm.Open(dialect, dbURI)
if err != nil{
	log.Fatal(err)
} else {
	fmt.Println("Successfully connected to database")
}



defer db.Close()



db.AutoMigrate(&models.User)

