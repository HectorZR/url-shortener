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
		g.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	if err := ValidateURL(url); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := ShortenURL(url)

	g.HTML(http.StatusCreated, "shortened-url.html", gin.H{"URL": shortURL})
}

func (c *Controller) RedirectURL(g *gin.Context) {
	shortURL := g.Param("shortURL")

	if shortURL == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Short URL is required"})
		return
	}

	originalUrl, err := GetOriginalURL(shortURL)

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	g.Redirect(http.StatusFound, originalUrl)
}
