package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) IndexView(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) ShortenURL(g *gin.Context) {
	url := g.PostForm("url")

	if url == "" {
		g.String(400, "URL is required")
		return
	}

	shortURL := ShortenURL(url)

	g.HTML(http.StatusCreated, "shortened-url.html", map[string]string{"URL": shortURL})
}
