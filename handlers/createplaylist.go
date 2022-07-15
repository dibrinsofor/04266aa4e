package handlers

import (
	"log"
	"net/http"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPlaylist(c *gin.Context) {
	var u models.Playlist

	if c.BindJSON(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	u.RandSlug = shortuuid.New()
	log.Print(u.RandSlug)

	u.ID = primitive.NewObjectID()
	err := models.AddUrlsToCollection(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":   "unable to store playlist.",
			"error_msg": err,
		})
	}

	// TODO sanitize dead or repeated links in queue
	c.JSON(http.StatusOK, gin.H{
		"message": "playlist successfully stored.",
		"data":    u.RandSlug,
		"links":   u.Urls,
	})
}
