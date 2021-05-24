package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(app *gin.Engine) {
	app.LoadHTMLGlob("webapp/*")
	app.GET("/login", Login)
	app.GET("/home", Home)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
