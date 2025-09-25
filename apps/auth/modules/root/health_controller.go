package root

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	GetHealth(c *gin.Context)
}

type healthController struct{}

func NewHealthController() HealthController {
	return &healthController{}
}

func (hc *healthController) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
