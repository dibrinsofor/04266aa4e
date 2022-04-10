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
	// TODO: setup auth before "/all" route
	r.GET("/health", handlers.HealthCheck())
	r.GET("/all", handlers.GetAllPlaylists)
	r.POST("/", handlers.AddUrl)
	r.GET("/:slug", handlers.GetUrls)

	r.Run(":8080")
}
