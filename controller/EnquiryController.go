package controller

import (
	"github.com/Shreetej/redwalls/models"
	"github.com/gin-gonic/gin"
)

func CreateEnquiry(c *gin.Context) {
	// Binding data in the body to the type Enquiry
	enquiry := models.Enquiry{}
	c.Bind(&enquiry)
	saveEnquiry := models.Enquiry{FirstName: enquiry.FirstName, LastName: enquiry.LastName, MobileNo: enquiry.MobileNo, EmailId: enquiry.EmailId, LandlineNo: enquiry.LandlineNo, EnquiryDetails: enquiry.EnquiryDetails}
	result:= models.DB.Create(&saveEnquiry)
	if result.Error != nil{
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{
		"enquiry":saveEnquiry,
	})
}

func UpdateEnquiry(c *gin.Context) {
	id := c.Param("id")
	enquiry := models.Enquiry{}
	c.Bind(&enquiry)

	var updateEnquiry models.Enquiry
	models.DB.First(&updateEnquiry,id)
	models.DB.Model(&updateEnquiry).Updates(models.Enquiry{FirstName: enquiry.FirstName,LastName: enquiry.LastName, MobileNo: enquiry.MobileNo,LandlineNo: enquiry.LandlineNo,EnquiryDetails: enquiry.EnquiryDetails, EmailId: enquiry.EmailId})
	c.JSON(200, gin.H{
		"enquiry":updateEnquiry,
	})
}

func GetAllEnquiries(c *gin.Context) {
	var enquiries []models.Enquiry
	models.DB.Find(&enquiries)
	c.JSON(200,gin.H{
		"enquries":enquiries,
	})
}

func GetEnquiry(c *gin.Context) {
	id := c.Param("id")
	var enquiry models.Enquiry
	models.DB.First(&enquiry,id)
	c.JSON(200,gin.H{
		"enquries":enquiry,
	})
}

func DeleteEnquiry(c *gin.Context) {
	id := c.Param("id")
	models.DB.Delete(&models.Enquiry{},id)
	c.Status(200)
}