package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Replace "your-connection-string" and "your-database" with your actual values
	initMongoDB("mongodb+srv://devanshv22:Devanshv17@cluster0.ccgq7vm.mongodb.net/?retryWrites=true&w=majority", "your-database")

	defer closeMongoDB()

	// Initialize Gin router
	router := gin.Default()

	// Handle requests to the root path
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, welcome to your application!")
	})

	// Set up authentication routes
	authRoutes(router)

	// Set up todo routes
	todoRoutes(router)

	// Run the server
	router.Run(":8080")
}
