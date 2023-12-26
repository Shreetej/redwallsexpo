package main

import (
	"fmt"

	"github.com/Shreetej/redwalls/models"
	"github.com/Shreetej/redwalls/routes"
) 

func init(){
	models.ConnectToDB()
	routes.Route()
	
}

func main() {
	
	fmt.Println("Welcome to Redwalls.")
}