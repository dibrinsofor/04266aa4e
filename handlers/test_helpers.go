package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTestPlaylist(t *testing.T) *models.Playlist {
	f := faker.New()

	slug := f.RandomStringWithLength(5)

	numberOfUrls := f.IntBetween(1, 8)
	var urls []string
	n := 1
	for n < numberOfUrls {
		n += 1
		urls = append(urls, f.Internet().URL())
	}

	testPlaylist := models.Playlist{
		ID:       primitive.NewObjectID(),
		Urls:     urls,
		RandSlug: slug,
	}

	return &testPlaylist
}

func SeedDB(playlist *models.Playlist) {
	err := godotenv.Load("../.env")
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

	defer client.Disconnect(ctx)
	_, err = client.Database("urlplaylists").Collection("urlplaylists_test").InsertOne(ctx, playlist)
	if err != nil {
		log.Printf("Unable to persist test playlist to database: %v", err)
	}
}

func MakeTestRequest(t *testing.T, method string, route string, body interface{}) *http.Request {
	requestBody, err := json.Marshal(body)
	assert.NoError(t, err)

	method = strings.ToUpper(method)

	request, err := http.NewRequest(method, route, bytes.NewReader(requestBody))
	assert.NoError(t, err)

	return request
}

func DecodeResponse(t *testing.T, response *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &responseBody))
	return responseBody
}
