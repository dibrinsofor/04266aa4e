package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 10
	connectionStringTemplate = "mongodb+srv://%s:%s@cluster0.r3cqf.mongodb.net/%s?retryWrites=true&w=majority"
)

func ConstructURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	dbName := os.Getenv("MONGODB_DBNAME")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, dbName)
	return connectionURI
}

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	client, err := mongo.NewClient(options.Client().ApplyURI(ConstructURI()))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to create mongodb cluster: %v", err)
	}

	// Pinging to make sure we can connect to db
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping mongodb cluster: %v", err)
	}
	fmt.Print("Connected to MongoDB")
	return client, ctx, cancel
}

func AddUrlsToCollection(playlist *Playlist) error {
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	playlist.ID = primitive.NewObjectID()

	_, err := client.Database("urlplaylists").Collection("urlplaylists").InsertOne(ctx, playlist)
	if err != nil {
		log.Printf("Unable to persist playlist to database: %v", err)
		return err
	}

	return nil
}
