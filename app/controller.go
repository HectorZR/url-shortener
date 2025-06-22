package app

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) IndexView(g *gin.Context) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		g.String(500, "Error parsing template")
		return
	}

	g.Status(http.StatusOK)
	t.Execute(g.Writer, nil)
}

func (c *Controller) ShortenURL(g *gin.Context) {
	url := g.PostForm("url")

	if url == "" {
		g.String(400, "URL is required")
		return
	}

	shortURL := ShortenURL(url)

	t, err := template.ParseFiles("templates/shortened-url.html")

	if err != nil {
		fmt.Println(err)
		g.String(500, "Error parsing template")
		return
	}

	g.Status(http.StatusCreated)

	t.Execute(g.Writer, map[string]string{"URL": shortURL})
}
