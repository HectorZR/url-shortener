package shortener

import (
	"github.com/gin-gonic/gin"
)

type ShortenerRoutes struct{}

func (sr *ShortenerRoutes) LoadRoutes(server *gin.Engine) {
	controller := &Controller{}

	server.GET("/", controller.IndexView)
	server.POST("/shorten", controller.ShortenURL)
	server.GET("/x/:shortURL", controller.RedirectURL)
}
