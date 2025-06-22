package app

import (
	"github.com/gin-gonic/gin"
)

type AppRoutes struct{}

func (r *AppRoutes) InitAppRoutes(server *gin.Engine) {
	controller := &Controller{}

	server.GET("/", controller.IndexView)

	server.POST("/shorten", controller.ShortenURL)
}
