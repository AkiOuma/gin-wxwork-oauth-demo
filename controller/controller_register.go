package controller

import "github.com/gin-gonic/gin"

func RegisterController(app *gin.Engine) {
	OAuthController(app)
	HomeController(app)
}
