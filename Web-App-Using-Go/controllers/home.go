// controllers/home.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Home Page",
		"message": "Welcome to the Simple Web Server!",
	})
}

func ContactForm(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", nil)
}

func ContactSubmit(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	if name == "" || email == "" || message == "" {
		c.HTML(http.StatusBadRequest, "contact.tmpl", gin.H{
			"error": "All fields are required.",
		})
		return
	}

	// Process the data (e.g., save to database)
	// For this example, we'll just render a thank-you page.

	c.HTML(http.StatusOK, "thankyou.html", gin.H{
		"name": name,
	})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Submit(c *gin.Context) {
	var json struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Process the data (e.g., save to database)

	c.JSON(http.StatusOK, gin.H{
		"status": "submission received",
		"name":   json.Name,
		"email":  json.Email,
	})
}
