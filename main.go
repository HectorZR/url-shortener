package main

import (
	"hectorzurga.com/url-shortener/app"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	var appRoutes app.AppRoutes

	appRoutes.InitAppRoutes(server)

	server.Run(":8000")
}
