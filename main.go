package main

import (
	"fmt"
	"oauth/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	controller.RegisterController(app)
	app.Run(":7001")
	fmt.Println("Oauth")
}
