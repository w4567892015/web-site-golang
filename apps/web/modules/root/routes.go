package root

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	healthController := NewHealthController()

	rootRoutes := server.Group("")
	{
		rootRoutes.GET("/health", healthController.GetHealth)
	}
}
