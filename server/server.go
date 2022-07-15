package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/dibrinsofor/urlplaylists/handlers"
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to retrieve user IP",
			})

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

func SetupServer() *gin.Engine {
	r := gin.Default()

	r.Use(Cors())
	r.Use(Limit())

	// TODO: setup auth before "/all" route
	r.GET("/health", handlers.HealthCheck())
	r.GET("/all", handlers.GetAllPlaylists)
	r.POST("/", handlers.AddPlaylist)
	r.GET("/:slug", handlers.GetUrls)

	return r
}
