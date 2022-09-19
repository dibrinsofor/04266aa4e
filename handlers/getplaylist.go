package handlers

import (
	"fmt"
	"net/http"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/gin-gonic/gin"
)

func GetUrls(c *gin.Context) {
	slug := c.Param("rand_slug")
	playlist, err := models.FindPlaylistBySlug(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to find playlist.",
		})
		return
	}
	fmt.Print(playlist)

	c.JSON(http.StatusOK, gin.H{
		"message":     "playlist retrieved succesfully",
		"title":       playlist.Title,
		"description": playlist.Description,
		"urls":        playlist.Urls,
	})
}
