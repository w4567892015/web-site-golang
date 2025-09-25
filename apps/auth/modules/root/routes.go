package root

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	healthController := NewHealthController()

	routes := server.Group("")
	{
		routes.GET("/health", healthController.GetHealth)
	}
}
