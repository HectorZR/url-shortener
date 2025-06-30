package main

import (
	"github.com/HectorZR/url-shortener/modules"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.StaticFile("static/utils.js", "./static/utils.js")

	modules.InitRoutes(server)

	server.Run(":8000")
}
