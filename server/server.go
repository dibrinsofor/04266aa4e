package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/dibrinsofor/urlplaylists/handlers"
	"github.com/dibrinsofor/urlplaylists/redis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	return cors.New(config)
}

// TODO abstract rate limiting middleware to alt file, also maybe not use IP
type User struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var users = make(map[string]*User)
var mu sync.Mutex

func InitBg() {
	go SanitizeUser()
}

func SanitizeUser() {
	for {
		time.Sleep(30 * time.Minute)

		mu.Lock()
		for ip, v := range users {
			if time.Since(v.lastSeen) > 72*time.Hour {
				delete(users, ip)
			}
		}
		mu.Unlock()
	}
}

func GetUserLimit(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	u, exists := users[ip]
	if !exists {
		limiter := rate.NewLimiter(1, 3)
		users[ip] = &User{limiter, time.Now()}

	}

	u.lastSeen = time.Now()
	return u.limiter
}

func Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if ip != "" {
			c.Next()
			// c.JSON(http.StatusInternalServerError, gin.H{
			// 	"message": "failed to retrieve user IP",
			// })

		}

		limiter := GetUserLimit(ip)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "slow down, that's too many requests",
			})
			return
		}
		c.Next()
	}
}

var Responses = []string{}

func Idempotent() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get idempotency key,
		id_key := c.Request.Header.Get("Idempotency-Key")
		IdemObject, err := redis.FindTransByIdemKey(id_key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "That's odd, try again later",
			})
		}
		if IdemObject == nil {
			c.Next()
		} else {
			switch IdemObject.StatusCode {
			case http.StatusOK:
				c.JSON(http.StatusOK, IdemObject.Response)
				// do this when the trans already succeeded
				// return map string interface of response?
			case http.StatusBadRequest:
				c.JSON(http.StatusBadRequest, IdemObject.Response)
				// do this when the client made a bad request the first time
			case http.StatusInternalServerError:
				c.Next()
				// do this when the trans failed the first time
			}
		}
		// if id_key == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"message": "missing header",
		// 	})
		// }
		// if object does not exist nko? send them to c.next()

		c.Next()
		// setup redis and search in-memory kv store for id_key
		// if key exists read it's corresponding action and do the needful
	}
}

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.Use(Cors())
	r.Use(Limit())

	// TODO: setup auth before "/all" route
	r.GET("/health", handlers.HealthCheck())
	r.GET("/all", handlers.GetAllPlaylists)
	r.POST("/", Idempotent(), handlers.AddPlaylist)
	r.GET("/:slug", handlers.GetUrls)

	return r
}
