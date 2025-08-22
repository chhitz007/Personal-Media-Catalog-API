package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"bookapi/config"
	"bookapi/routes"
)

func init() {
	// Load environment variables from .env file at the start
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 🟨 DEBUG: Print JWT_SECRET to verify it's loaded
	fmt.Println("JWT_SECRET (from .env):", os.Getenv("JWT_SECRET"))
}

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	routes.RegisterAuthRoutes(router)
	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

