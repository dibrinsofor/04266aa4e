package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dibrinsofor/urlplaylists/handlers"
	"github.com/dibrinsofor/urlplaylists/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddPlaylistSucceed(t *testing.T) {
	ts := server.SetupServer()

	samplePlaylists := handlers.CreateTestPlaylist(t)
	handlers.SeedDB(samplePlaylists)

	request := handlers.MakeTestRequest(t, "POST", "/", map[string]interface{}{
		"_id":       samplePlaylists.ID,
		"urls":      samplePlaylists.Urls,
		"rand_slug": samplePlaylists.RandSlug,
	})

	response := bootsrapRequest(request, ts)
	responseBody := handlers.DecodeResponse(t, response)
	assert.Equal(t, "playlist successfully stored.", responseBody["message"])
}

func bootsrapRequest(request *http.Request, routeHandlers *gin.Engine) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	routeHandlers.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
