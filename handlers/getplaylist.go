package handlers

import (
	"fmt"
	"net/http"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/gin-gonic/gin"
)

type RequestUrlPlaylist struct {
	RandSlug string `uri:"slug" binding:"required"`
}

func GetUrl(c *gin.Context) {
	var u RequestUrlPlaylist
	var playlist models.Playlist

	if c.ShouldBindUri(&u) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
	}
	slug := u.RandSlug
	ctx, cursor := models.FindPlaylistBySlug(slug)
	if err := cursor.All(ctx, &playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to find user.",
		})
	}
	defer cursor.Close(ctx)
	fmt.Print(playlist)

	// TODO check for invalid slug passed or too long
	// TODO find user by slug

	c.JSON(http.StatusOK, gin.H{
		"message": "playlist retrieved succesfully",
		"data":    playlist.Urls,
	})
}
