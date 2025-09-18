package main

import (
	"os"

	logger "libs/otel"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logger.Log("web", "Application started.")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:" + port)
}
