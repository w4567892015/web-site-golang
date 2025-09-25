package core

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/fosite/handler/openid"
)

type CoreController interface {
	authorizeEndpoint(c *gin.Context)
	tokenEndpoint(c *gin.Context)
	configuration(c *gin.Context)
}

type coreController struct{}

func NewCoreController() CoreController {
	return &coreController{}
}

func (cc *coreController) configuration(c *gin.Context) {
}

func (cc *coreController) authorizeEndpoint(c *gin.Context) {
	ar, err := OAuth2.NewAuthorizeRequest(c, c.Request)
	if err != nil {
		OAuth2.WriteAuthorizeError(c, c.Writer, ar, err)
		return
	}

	response, err := OAuth2.NewAuthorizeResponse(c, ar, &openid.DefaultSession{})
	if err != nil {
		OAuth2.WriteAuthorizeError(c, c.Writer, ar, err)
		return
	}

	OAuth2.WriteAuthorizeResponse(c, c.Writer, ar, response)
}

func (cc *coreController) tokenEndpoint(c *gin.Context) {
	mySessionData := &openid.DefaultSession{}

	tr, err := OAuth2.NewAccessRequest(c, c.Request, mySessionData)
	if err != nil {
		OAuth2.WriteAccessError(c, c.Writer, tr, err)
		return
	}

	// 產生並回應存取令牌
	response, err := OAuth2.NewAccessResponse(c, tr)
	if err != nil {
		OAuth2.WriteAccessError(c, c.Writer, tr, err)
		return
	}

	OAuth2.WriteAccessResponse(c, c.Writer, tr, response)
}
