package handlers

import (
	"net/http"
	"urlplaylists/models"

	"github.com/gin-gonic/gin"
)

func AddUrl(c *gin.Context) {
	var u models.Playlist

	if c.BindJSON(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	err := db.CreateURL(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "unable to store playlist.",
			"error_msg": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "playlist successfully stored.",
		"data":    u,
	})
}
