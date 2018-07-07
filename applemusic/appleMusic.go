package applemusic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../music"
)

// LinkHandler Represents a handler for requests against apple music
type LinkHandler struct {
}

// GetAlbum Fetches an album with the apple music identifier
func (linkHandler LinkHandler) GetAlbum(id string) (music.Response, error) {
	musicResponse := music.Response{}
	appleMusicResponse := AlbumResponse{}
	request, err := http.NewRequest("GET", "https://api.music.apple.com/v1/catalog/de/albums/"+id, nil)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ijg5N0E1RkE4WkEifQ.eyJpYXQiOjE1MzAxMDUwNjQsImV4cCI6MTU0NTY1NzA2NCwiaXNzIjoiWkY5OUdFOVI1VyJ9.JRN6e__NCO8Yjhj2ynJV20RbPOuNDo9WLcR_lYg1B348ea4BembEqraV53MF-c14jxKYk_0pRjjJlhmF3lkmdw")
	if err != nil {
		return musicResponse, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return musicResponse, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return musicResponse, err
	}

	err = json.Unmarshal(data, &appleMusicResponse)
	if err != nil {
		return musicResponse, err
	}

	musicResponse.MediaType = "album"
	musicResponse.Artist = appleMusicResponse.Data[0].Attributes.ArtistName
	musicResponse.Album = appleMusicResponse.Data[0].Attributes.Name
	return musicResponse, nil
}

// GetArtist Fetches informations for an apple music artist
func (linkHandler LinkHandler) GetArtist(id string) (music.Response, error) {
	musicResponse := music.Response{}
	appleMusicResponse := ArtistResponse{}
	request, err := http.NewRequest("GET", "https://api.music.apple.com/v1/catalog/de/artists/"+id, nil)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ijg5N0E1RkE4WkEifQ.eyJpYXQiOjE1MzAxMDUwNjQsImV4cCI6MTU0NTY1NzA2NCwiaXNzIjoiWkY5OUdFOVI1VyJ9.JRN6e__NCO8Yjhj2ynJV20RbPOuNDo9WLcR_lYg1B348ea4BembEqraV53MF-c14jxKYk_0pRjjJlhmF3lkmdw")
	if err != nil {
		return musicResponse, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return musicResponse, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return musicResponse, err
	}

	err = json.Unmarshal(data, &appleMusicResponse)
	if err != nil {
		return musicResponse, err
	}

	musicResponse.MediaType = "artist"
	musicResponse.Artist = appleMusicResponse.Data[0].Attributes.Name
	return musicResponse, nil
}

// GetSong Fetches informations for a apple music song
func (linkHandler LinkHandler) GetSong(id string) (music.Response, error) {
	musicResponse := music.Response{}
	appleMusicResponse := SongResponse{}
	request, err := http.NewRequest("GET", "https://api.music.apple.com/v1/catalog/de/songs/"+id, nil)
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ijg5N0E1RkE4WkEifQ.eyJpYXQiOjE1MzAxMDUwNjQsImV4cCI6MTU0NTY1NzA2NCwiaXNzIjoiWkY5OUdFOVI1VyJ9.JRN6e__NCO8Yjhj2ynJV20RbPOuNDo9WLcR_lYg1B348ea4BembEqraV53MF-c14jxKYk_0pRjjJlhmF3lkmdw")
	if err != nil {
		return musicResponse, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return musicResponse, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return musicResponse, err
	}

	err = json.Unmarshal(data, &appleMusicResponse)
	if err != nil {
		return musicResponse, err
	}

	musicResponse.MediaType = "song"
	musicResponse.Artist = appleMusicResponse.Data[0].Attributes.ArtistName
	musicResponse.Album = appleMusicResponse.Data[0].Attributes.AlbumName
	musicResponse.Song = appleMusicResponse.Data[0].Attributes.Name
	return musicResponse, nil
}
