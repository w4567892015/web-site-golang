package main

import (
	"embed"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"apps/auth/middlewares"
	"apps/auth/modules/core"
	"apps/auth/modules/login"
	"apps/auth/modules/root"
)

// authorizeHandler 現在被拆分為 GET 和 POST 兩個處理器
// func authorizeGetHandler(c *gin.Context, oauth2 fosite.OAuth2Provider) {
// 	ctx := c.Request.Context()
// 	ar, err := oauth2.NewAuthorizeRequest(ctx, c.Request)
// 	if err != nil {
// 		log.Print("Error occurred in NewAuthorizeRequest: %+v", err)
// 		oauth2.WriteAuthorizeError(ctx, c.Writer, ar, err)
// 		return
// 	}

// 	// 顯示同意授權頁面
// 	html := `
// 		<h1>同意授權</h1>
// 		<p>應用程式 <stron` + ar.GetClient().GetID() + `</strong> 想要存取您的以下資訊：<strong>` + strings.Join(ar.GetRequestedScopes(), ", ") + `</strong></p>
// 		<form metho"post">
// 			<inptype="hidden" name="state" value="` + ar.GetState() + `"/>
// 			<inptype="submit" value="允許 (Allow)" />
// 		</form>
// `
// 	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
// }

// // tokenHandler 的邏輯不變，只是改為 Gin 的形式
// func tokenHandler(c *gin.Context, oauth2 fosite.OAuth2Provider) {
// 	ctx := c.Request.Context()
// 	session := &openid.DefaultSession{}

// 	ar, err := oauth2.NewAccessRequest(ctx, c.Request, session)
// 	if err != nil {
// 		log.Print("Error occurred in NewAccessRequest: %+v", err)
// 		oauth2.WriteAccessError(ctx, c.Writer, ar, err)
// 		return
// 	}

// 	if ar.GetGrantTypes().ExactOne("refresh_token") {
// 		ar.GrantScop("openid")
// 	}

// 	response, err := oauth2.NewAccessResponse(ctx, ar)
// 	if err != nil {
// 		log.Print("Error occurred in NewAccessResponse: %+v", err)
// 		oauth2.WriteAccessError(ctx, c.Writer, ar, err)
// 		return
// 	}

// 	oauth2.WriteAccessResponse(ctx, c.Writer, ar, response)
// }

func run(server *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := server.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

//go:embed view/*
var embeddedTemplates embed.FS

func main() {
	server := gin.Default()

	server.SetHTMLTemplate(
		template.Must(
			template.ParseFS(embeddedTemplates, "view/*.html"),
		),
	)

	server.Use(middlewares.CORSMiddleware())

	// Register module routes
	root.RegisterRoutes(server)
	login.RegisterRoutes(server)
	core.RegisterRoutes(server)

	run(server)
}
