package spotify

import (
	"../music"
)

// LinkHandler Represents a handler for requests against apple music
type LinkHandler struct {
}

// GetArtist Fetches an album with the spotify identifier
func (linkHandler LinkHandler) GetArtist(id string) music.Response {
	response := music.Response{
		MediaType: "artist",
		Artist:    "test"}
	return response
}

// GetAlbum Fetches an album with the spotify identifier
func (linkHandler LinkHandler) GetAlbum(id string) music.Response {
	response := music.Response{
		MediaType: "album",
		Artist:    "test",
		Album:     "test"}
	return response
}

// GetSong Fetches informations for a spotify song
func (linkHandler LinkHandler) GetSong(id string) music.Response {
	response := music.Response{
		MediaType: "song",
		Artist:    "test",
		Album:     "test",
		Song:      "test"}
	return response
}
