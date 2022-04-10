package handlers

import (
	"net/http"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/gin-gonic/gin"
)

func GetAllPlaylists(c *gin.Context) {

	playlists, err := models.FindAllPlaylists()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to find all playlists.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": playlists,
	})
}
