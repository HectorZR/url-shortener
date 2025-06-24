package main

import (
	"github.com/HectorZR/url-shortener/app"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.StaticFile("static/utils.js", "./static/utils.js")

	app.InitAppRoutes(server)

	server.Run(":8000")
}
