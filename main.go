package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		t, err := template.ParseFiles("templates/index.html")

		if err != nil {
			c.String(500, "Error parsing template")
			return
		}

		c.Status(http.StatusOK)
		t.Execute(c.Writer, nil)
	})
	router.Run(":8000")
}
