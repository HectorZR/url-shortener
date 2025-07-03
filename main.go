package main

import (
	"fmt"

	"github.com/HectorZR/url-shortener/modules"
	"github.com/HectorZR/url-shortener/shared"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.StaticFile("static/utils.js", "./static/utils.js")

	modules.InitRoutes(server)

	server.Run(fmt.Sprintf(":%s", shared.GetEnvVars()["PORT"]))
}
