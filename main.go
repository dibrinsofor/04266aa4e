package main

import (
	"github.com/dibrinsofor/urlplaylists/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handlers.HealthCheck())
	r.POST("/list", handlers.AddUrl)

	r.Run("localhost:8080")
}
