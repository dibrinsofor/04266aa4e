package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/dibrinsofor/urlplaylists/redis"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPlaylist(c *gin.Context) {
	var u models.Playlist
	var i models.IdemKey

	i.ID = c.Request.Header.Get("Idempotency-Key")

	if c.BindJSON(&u) != nil {
		i.Response = map[string]interface{}{
			"message": "failed to parse user request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		}
		i.StatusCode = http.StatusBadRequest
		i.CreatedAt = time.Now().Format("2017-09-07 2:3:5 PM")
		c.JSON(http.StatusBadRequest, i.Response)
		return
	}

	u.RandSlug = shortuuid.New()
	log.Print(u.RandSlug)

	u.ID = primitive.NewObjectID()
	err := models.AddUrlsToCollection(&u)
	if err != nil {
		i.Response = map[string]interface{}{
			"message":   "unable to store playlist.",
			"error_msg": err,
		}
		i.StatusCode = http.StatusInternalServerError
		i.CreatedAt = time.Now().Format("2017-09-07 2:3:5 PM")
		c.JSON(http.StatusInternalServerError, i.Response)
	}

	i.StatusCode = http.StatusOK
	i.Response = map[string]interface{}{
		"message": "playlist successfully stored.",
		"data":    u.RandSlug,
		"links":   u.Urls,
	}
	i.CreatedAt = time.Now().Format("2017-09-07 2:3:5 PM")
	_, err = redis.AddIdemKey(&i)
	if err != nil {
		i.Response = map[string]interface{}{
			"message":   "unable to store playlist.",
			"error_msg": err,
		}
		i.StatusCode = http.StatusInternalServerError
		i.CreatedAt = time.Now().Format("2017-09-07 2:3:5 PM")
		c.JSON(http.StatusInternalServerError, i.Response)
	}

	// TODO sanitize dead or repeated links in queue
	c.JSON(http.StatusOK, i.Response)
}
