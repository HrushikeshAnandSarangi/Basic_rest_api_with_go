package main

import (
	"log"
	"os"

	"github.com/HrushikeshAnandSarangi/go-rest/config"
	"github.com/HrushikeshAnandSarangi/go-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.AuthRoutes(r)

	// Get port from environment (Render sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // fallback for local
	}

	log.Printf("Listening on port %s...\n", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
