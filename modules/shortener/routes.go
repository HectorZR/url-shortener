package shortener

import (
	"github.com/gin-gonic/gin"
)

type ShortenerRoutes struct{}

func (sr *ShortenerRoutes) LoadRoutes(server *gin.Engine) {
	controller := &Controller{}

	group := server.Group("/shortener")
	group.GET("/", controller.IndexView)
	group.POST("/shorten", controller.ShortenURL)
	group.GET("/x/:shortURL", controller.RedirectURL)
}
