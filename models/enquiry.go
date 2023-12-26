package models

import "gorm.io/gorm"

type Enquiry struct {
	gorm.Model
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	MobileNo       uint   `json:"mobileNo"`
	EmailId		   string `json:"emailId"`
	LandlineNo     string `json:"landlineNo"`
	EnquiryDetails string `json:"enquiryDetails"`
}