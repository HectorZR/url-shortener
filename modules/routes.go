package modules

import (
	"github.com/gin-gonic/gin"

	"github.com/HectorZR/url-shortener/modules/shortener"
)

type IRoute interface {
	LoadRoutes(server *gin.Engine)
}

func InitRoutes(server *gin.Engine) {
	routes := []IRoute{
		&shortener.ShortenerRoutes{},
	}

	for _, route := range routes {
		route.LoadRoutes(server)
	}
}
