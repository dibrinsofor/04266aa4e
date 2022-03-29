package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// (*mongo.Client, context.Context, context.CancelFunc)

func GetConnection() (*mongo.Client, context.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
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
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Printf("Failed to make db request: %s", err)
	// }
	// fmt.Println(databases)
	return client, ctx
}

func AddUrlsToCollection(playlist *Playlist) error {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	// playlist.ID = primitive.NewObjectID()

	testtt, err := client.Database("urlplaylists").Collection("urlplaylists").InsertOne(ctx, playlist)
	if err != nil {
		log.Printf("Unable to persist playlist to database: %v", err)
		return err
	}

	fmt.Println(testtt.InsertedID)

	return nil
}

func CheckPlaylistSlug(slug string) {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	filter := bson.D{{"rand_slug", slug}}

	var playlist bson.M
	err := client.Database("urlplaylists").Collection("urlplaylists").FindOne(ctx, filter).Decode(&playlist)
	if err != nil {

	}

}
