package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestUrlPlaylist struct {
	RandSlug string `json:"rand_slug,omitempty" bson:"rand_slug,omitempty"`
}

func GetUrl(c *gin.Context) {
	var u RequestUrlPlaylist

	if err

	// TODO find user by slug

	c.JSON(http.StatusOK, gin.H{
		"message": "playlist retrieved succesfully",
		"data":    u,
	})
}
