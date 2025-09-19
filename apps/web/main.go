package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"apps/web/middlewares"
	logger "libs/otel"

	"apps/web/modules/index"
	"apps/web/modules/root"
)

func run(server *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := server.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {

	logger.Log("web", "Application started.")

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	// Register module routes
	root.RegisterRoutes(server)
	index.RegisterRoutes(server)

	run(server)
}
