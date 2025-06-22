package main

import (
	"url-shortener/app"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	app.InitDB()

	app.InitAppRoutes(server)

	server.Run(":8000")
}
