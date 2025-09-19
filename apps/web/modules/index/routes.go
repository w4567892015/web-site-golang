package index

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	indexService := NewIndexService()
	indexController := NewIndexController(indexService)

	indexRoutes := server.Group("/api")
	{
		indexRoutes.GET("", indexController.GetIndex)
	}
}
