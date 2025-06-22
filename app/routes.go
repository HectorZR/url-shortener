package app

import (
	"github.com/gin-gonic/gin"
)

func InitAppRoutes(server *gin.Engine) {
	controller := &Controller{}

	server.GET("/", controller.IndexView)

	server.POST("/shorten", controller.ShortenURL)
}
