package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string 	`json:"userName"`
	Password  string 	`json:"-"`
	EmailId   string 	`json:"emailId"`
	FirstName string 	`json:"firstName"`
	LastName  string	`json:"lastName"`
}