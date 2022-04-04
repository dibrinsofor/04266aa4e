package handlers

import (
	"log"
	"net/http"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUrl(c *gin.Context) {
	var u models.Playlist

	if c.BindJSON(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	u.RandSlug = models.GetPlaylistSlug()
	log.Print(u.RandSlug)

	u.ID = primitive.NewObjectID()
	err := models.AddUrlsToCollection(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "unable to store playlist.",
			"error_msg": err,
		})
	}

	// TODO construct a new uri an pass that in the response too

	c.JSON(http.StatusOK, gin.H{
		"message": "playlist successfully stored.",
		"data":    u,
	})
}
