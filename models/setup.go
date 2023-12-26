package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(){

	var err error
	url := "host=localhost user=postgres password=roger dbname=redwalls port=5432 sslmode=disable"
	fmt.Println("In URL",url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil{
		panic("Connection to DB failed.")
	}
 
	//Migrate the schema
	db.AutoMigrate(&User{},&Enquiry{})
	DB = db
}