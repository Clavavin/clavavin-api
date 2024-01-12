package main

import (
	"fmt"

	"github.com/Clavavin/clavavin-api/controllers"
	"github.com/Clavavin/clavavin-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")

	// Connect to database
	database.DatabaseConnection()

	// Define router and routes
	r := gin.Default()
	r.POST("/wines", controllers.CreateWine)
	r.GET("/wines/:id", controllers.ReadWine)
	r.GET("/wines", controllers.ReadWines)
	r.PUT("/wines/:id", controllers.UpdateWine)
	r.DELETE("/wines/:id", controllers.DeleteWine)

	// Run the application on port 8080
	r.Run(":8080")
}
