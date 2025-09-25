package core

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	coreController := NewCoreController()

	routes := server.Group("/oauth2")
	{
		routes.GET("/.well-known/openid-configuration", coreController.configuration)
		routes.GET("/auth", coreController.authorizeEndpoint)
		// routes.POST("/authorize", coreController.authorizeEndpoint)
		routes.POST("/token", coreController.tokenEndpoint)
	}
}
