package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	indexService IndexService
}

func NewIndexController(s IndexService) *IndexController {
	return &IndexController{indexService: s}
}

func (ic *IndexController) GetIndex(ctx *gin.Context) {
	ctx.String(http.StatusOK, ic.indexService.GetHello())
}
