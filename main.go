package main

import (
	"log"
	"os"
	"url-shortener/models"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title URL Shortener
// @description Generate short link
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.ConnectDB()

	port := os.Getenv("SV_PORT")
	r := gin.Default()
	routes.Config(r)

	r.Run(":" + port)
}
