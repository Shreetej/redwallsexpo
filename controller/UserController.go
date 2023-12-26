package controller

import (
	"fmt"

	"github.com/Shreetej/redwalls/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct{
		UserName  string
		Password  string 
		EmailId   string
		FirstName string
		LastName  string
	}
	c.Bind(&body)
	fmt.Println(body)
	// tempUser := models.User{UserName: "Roger",Password: "Roger",FirstName: "Roger",LastName: "Federer", EmailId: "rghricks@gmail.com"}
	user :=models.User{UserName: body.UserName, Password: body.Password, EmailId: body.EmailId, FirstName: body.FirstName, LastName: body.LastName}
	result:= models.DB.Create(&user)
	if result.Error != nil{
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{
		"user":result,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	c.Bind(&user)

	var updateUser models.User
	models.DB.First(&updateUser,id)
	models.DB.Model(&updateUser).Updates(models.User{FirstName: user.FirstName,LastName: user.LastName, UserName: user.UserName,Password: user.Password,EmailId: user.EmailId})
	c.JSON(200, gin.H{
		"user":updateUser,
	})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(200,gin.H{
		"users":users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	models.DB.First(&user,id)
	c.JSON(200,gin.H{
		"user":user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	models.DB.Delete(&models.User{},id)
	c.Status(200)
}