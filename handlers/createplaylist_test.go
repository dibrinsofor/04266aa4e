package handlers_test

import (
	"testing"

	"github.com/dibrinsofor/urlplaylists/handlers"
	"github.com/dibrinsofor/urlplaylists/models"
)

func TestAddPlaylist(t *testing.T) *models.Playlist {
	samplePlaylists := handlers.CreateTestPlaylist(t)
	// Figure out wtf seed db is
	return samplePlaylists
}
