package server

import (
	"github.com/dibrinsofor/urlplaylists/handlers"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()

	// TODO: setup auth before "/all" route
	r.GET("/health", handlers.HealthCheck())
	r.GET("/all", handlers.GetAllPlaylists)
	r.POST("/", handlers.AddPlaylist)
	r.GET("/:slug", handlers.GetUrls)

	return r
}
