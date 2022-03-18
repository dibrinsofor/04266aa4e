package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
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

func CreateURL(playlist *Playlist) error {
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
