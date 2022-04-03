package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dibrinsofor/urlplaylists/lib"
	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// const (
// 	collection = client.Database("urlplaylists").Collection("urlplaylists")
// 	database   = client.Database("urlplaylists")
// )

func GetConnection() (*mongo.Client, context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Printf("Failed to create client: %s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to database cluster: %s", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Failed to ping mongodb cluster: %s", err)
	}
	return client, ctx
}

func AddUrlsToCollection(playlist *Playlist) error {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)
	result, err := client.Database("urlplaylists").Collection("urlplaylists").InsertOne(ctx, playlist)
	if err != nil {
		log.Printf("Unable to persist playlist to database: %v", err)
		return err
	}

	fmt.Println(result.InsertedID)

	return nil
}
func IsUniqueslug(slug string) error {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	var playlist models.Playlist
	filter := models.Playlist{Key: "rand_slug", Value: slug}
	count, err := client.Database("urlplaylists").Collection("urlplaylists").CountDocuments(ctx, filter)
	if err != nil {
		log.Print(err)
	}
	return err
}

// while slug is not unique, generate again

func GetPlaylistSlug() string {
	slug := lib.GenShortSlug()
	// err := IsUniqueslug(slug)
	// if err != nil {
	// 	log.Print(err)
	// }
	return slug
}
