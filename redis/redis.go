package redis

import (
	"crypto/tls"
	"os"
	"time"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
)

func ConnectRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "127.0.0.1:6379"
	}
	pass := os.Getenv("REDIS_PASSWORD")
	if pass == "" {
		pass = "password"
	}

	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
	rdb := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     pass,
		DB:           0,
		TLSConfig:    tlsConfig,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	return rdb
}

func AddIdemKey(user *models.IdemKey) ([]redis.Cmder, error) {
	rdb := ConnectRedis()

	val, err := rdb.Pipelined(func(rdb redis.Pipeliner) error {
		rdb.HSet(user.ID, "StatusCode", user.StatusCode)
		rdb.HSet(user.ID, "Response", user.Response)
		rdb.HSet(user.ID, "CreatedAt", user.CreatedAt)
		return nil
	})
	if err != nil {
		// panic(err)
	}

	defer rdb.Close()
	return val, nil
}

func FindTransByIdemKey(ID string) (*models.IdemKey, error) {
	rdb := ConnectRedis()

	IdemObject := &models.IdemKey{}
	val := rdb.HGetAll(ID).Val()
	err := mapstructure.Decode(val, &IdemObject)
	if err != nil {
		panic(err)
	}

	defer rdb.Close()
	return IdemObject, nil
}
