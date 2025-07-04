package shortener

import (
	"fmt"
	"net/http"

	"github.com/HectorZR/url-shortener/shared"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) IndexView(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", gin.H{
		"Path": shared.GetEnvVars().PathPrefix,
	})
}

func (c *Controller) ShortenURL(g *gin.Context) {
	url := g.PostForm("url")

	if err := ValidateURL(url); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURLEntity := ShortenURL(url, shared.InitDB())
	shortCode := shared.EncodeBase62(shortURLEntity.ID)

	shortUrl := fmt.Sprintf("%s%s/x/%s", g.Request.Host, shared.GetEnvVars().PathPrefix, shortCode)
	g.HTML(http.StatusCreated, "shortened-url.html", gin.H{"URL": shortUrl})
}

func (c *Controller) RedirectURL(g *gin.Context) {
	shortCode := g.Param("shortURL")

	if shortCode == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Short URL is required"})
		return
	}

	urlEntity, err := GetOriginalURL(shortCode, shared.InitDB())

	if err != nil {
		g.HTML(http.StatusNotFound, "not-found.html", nil)
		return
	}

	g.Redirect(http.StatusFound, urlEntity.OriginalURL)
}
