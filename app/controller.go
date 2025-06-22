package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) IndexView(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) ShortenURL(g *gin.Context) {
	url := g.PostForm("url")

	if err := ValidateURL(url); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURLEntity := ShortenURL(url, initDB())

	shortUrl := fmt.Sprint(g.Request.Host, "/x/", shortURLEntity.ShortURL)
	g.HTML(http.StatusCreated, "shortened-url.html", gin.H{"URL": shortUrl})
}

func (c *Controller) RedirectURL(g *gin.Context) {
	shortURL := g.Param("shortURL")

	if shortURL == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Short URL is required"})
		return
	}

	urlEntity, err := GetOriginalURL(shortURL, initDB())

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	g.Redirect(http.StatusFound, urlEntity.OriginalURL)
}
