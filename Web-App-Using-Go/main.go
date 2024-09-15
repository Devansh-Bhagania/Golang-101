// main.go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/devansh_bhagania/simple-web-server/controllers"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static", "./static")

	// Apply custom middleware
	// router.Use(middleware.AuthMiddleware())

	// Define routes
	router.GET("/", controllers.HomePage)
	router.GET("/contact", controllers.ContactForm)
	router.POST("/contact", controllers.ContactSubmit)

	// Grouping routes
	api := router.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
		api.POST("/submit", controllers.Submit)
	}

	// Start the server on port 8080
	router.Run(":8080")
}
