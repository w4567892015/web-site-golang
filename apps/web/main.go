package main

import (
	"context"
	"fmt"
	"libs/otel"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"apps/web/middlewares"
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
	ctx := context.Background()
	telem, err := otel.NewTelemetry(ctx)
	if err != nil {
		fmt.Println("failed to create telemetry, falling back to no-op telemetry")
	}
	defer telem.Shutdown(ctx)

	server := gin.Default()
	server.Use(otel.Middlewares())
	server.Use(middlewares.CORSMiddleware())

	// Register module routes
	root.RegisterRoutes(server)
	index.RegisterRoutes(server)

	run(server)
}
