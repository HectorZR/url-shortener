package main

import (
	"fmt"

	"github.com/HectorZR/url-shortener/modules"
	"github.com/HectorZR/url-shortener/shared"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	envVars := shared.GetEnvVars()

	if envVars.Env == shared.PROD {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.Static("/static", "./static")

	modules.InitRoutes(server)

	port := fmt.Sprintf(":%s", envVars.Port)

	if envVars.Env == shared.PROD {
		fmt.Printf("App running on port %s\n", port)
	}

	server.Run(port)
}
