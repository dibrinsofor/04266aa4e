package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dibrinsofor/urlplaylists/models"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartMockServer(request *http.Request, routeHandlers *gin.Engine) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	routeHandlers.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

func CreateTestPlaylist(t *testing.T) *models.Playlist {
	f := faker.New()

	slug := f.RandomStringWithLength(5)

	numberOfUrls := f.IntBetween(1, 8)
	var urls []string
	n := 1
	for n < numberOfUrls {
		n += 1
		urls = append(urls, f.Internet.Urls)
	}

	testPlaylist := models.Playlist{
		ID:       primitive.NewObjectID(),
		Urls:     urls,
		RandSlug: slug,
	}

	return &testPlaylist

}

func MakeTestRequest(t *testing.T, method string, route string, body interface{}) *http.Request {
	requestBody, err := json.Marshal(body)
	assert.NoError(t, err)

	method = strings.ToUpper(method)

	request, err := http.NewRequest(method, route, bytes.NewReader(requestBody))
	assert.NoError(t, err)

	return request
}
