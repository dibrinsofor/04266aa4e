package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dibrinsofor/urlplaylists/lib"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func GetConnection() (*mongo.Client, context.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		// log.Fatal("Error loading .env file")
		log.Fatalf("Error: %s", err)
	}

	client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Printf("Failed to create client: %s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*50)
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
func IsUniqueslug(slug string) bool {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	filter := bson.D{{Key: "rand_slug", Value: slug}}
	count, err := client.Database("urlplaylists").Collection("urlplaylists").CountDocuments(ctx, filter)
	if err != nil {
		log.Print(err)
	}
	if count != 0 {
		log.Print("slug not unique")
		return false
	}
	return true
}

// TODO: consider storing hash of slug instead of wait generation
func GetPlaylistSlug() string {
	slug := lib.GenShortSlug()

	for !IsUniqueslug(slug) {
		slug = lib.GenShortSlug()
	}
	return slug
}

// TODO: figure out indexing for slug
func FindPlaylistBySlug(slug string) (Playlist, error) {
	client, ctx := GetConnection()

	var playlist Playlist

	err := client.Database("urlplaylists").Collection("urlplaylists").FindOne(ctx, Playlist{RandSlug: slug}).Decode(&playlist)
	if err != nil {
		log.Fatal(err)
	}
	return playlist, err
}

func FindAllPlaylists() (*[]Playlist, error) {
	client, ctx := GetConnection()
	cursor, err := client.Database("urlplaylists").Collection("urlplaylists").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var playlists []Playlist
	if err = cursor.All(ctx, &playlists); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &playlists, nil
}
