package main

import (
	"github.com/dibrinsofor/urlplaylists/handlers"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{""}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))
	r.GET("/", handlers.HealthCheck())
	r.POST("/list", handlers.AddUrl)

	r.Run("localhost:8080")
}
