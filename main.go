package main

import (
	"github.com/dibrinsofor/urlplaylists/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{""}
	// corsConfig.AllowCredentials = true
	// corsConfig.AddAllowMethods("OPTIONS")

	// r.Use(cors.New(corsConfig))
	r.GET("/health", handlers.HealthCheck())
	r.POST("/", handlers.AddUrl)

	r.Run(":8080")
}
