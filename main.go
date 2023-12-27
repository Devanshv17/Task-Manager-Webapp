package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB
	initMongoDB()
	defer closeMongoDB()

	r := gin.Default()

	// Routes
	authRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s...\n", port)
	log.Fatal(r.Run(":" + port))
}
