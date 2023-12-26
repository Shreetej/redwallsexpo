package routes

import (
	"github.com/Shreetej/redwalls/controller"
	"github.com/gin-gonic/gin"
)
 

func Route()  {
	r := gin.Default()

	// User Routes
	r.POST("/users/",controller.CreateUser)
	r.GET("/users/",controller.GetAllUsers)
	r.PUT("/users/:id",controller.UpdateUser)
	r.DELETE("/users/:id",controller.DeleteUser)
	r.GET("/users/:id",controller.GetUser)

	// Enquiry Routes
	r.POST("/enquiry/",controller.CreateEnquiry)
	r.GET("/enquiry/",controller.GetAllEnquiries)
	r.PUT("/enquiry/:id",controller.UpdateEnquiry)
	r.DELETE("/enquiry/:id",controller.DeleteEnquiry)
	r.GET("/enquiry/:id",controller.GetEnquiry)

	r.Run()
}