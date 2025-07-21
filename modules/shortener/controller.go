package shortener

import (
	"fmt"
	"net/http"

	"github.com/HectorZR/url-shortener/shared"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) IndexView(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", gin.H{"SiteKey": shared.GetEnvVars().SiteKey})
}

func (c *Controller) ShortenURL(g *gin.Context) {
	url := g.PostForm("url")
	recaptchaToken := g.PostForm("g-recaptcha-response")
	env := shared.GetEnvVars()

	// Validate reCAPTCHA with Google
	if err := validateRecaptcha(env.ProjectID, env.SiteKey, recaptchaToken, env.CaptchaAction); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate URL format
	if err := ValidateURL(url); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process URL shortening
	shortURLEntity := ShortenURL(url, shared.InitDB())
	shortCode := shared.EncodeBase62(shortURLEntity.ID)

	protocol := "https"
	shortUrl := fmt.Sprintf("%s://%s/x/%s", protocol, g.Request.Host, shortCode)
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
