package applemusic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"../music"
)

// LinkHandler Represents a handler for requests against apple music
type LinkHandler struct {
}

func (linkHandler LinkHandler) Search(information music.Information) (string, error) {
	types := linkHandler.mapMediaType(information.MediaType)
	searchTerm := linkHandler.getSearchTerm(information)
	request, err := http.NewRequest("GET", "https://api.music.apple.com/v1/catalog/de/search?term="+searchTerm+"&types="+types, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ijg5N0E1RkE4WkEifQ.eyJpYXQiOjE1MzAxMDUwNjQsImV4cCI6MTU0NTY1NzA2NCwiaXNzIjoiWkY5OUdFOVI1VyJ9.JRN6e__NCO8Yjhj2ynJV20RbPOuNDo9WLcR_lYg1B348ea4BembEqraV53MF-c14jxKYk_0pRjjJlhmF3lkmdw")

	client := http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer httpResponse.Body.Close()

	data, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	searchResponse := SearchResponse{}
	err = json.Unmarshal(data, &searchResponse)
	if err != nil {
		return "", err
	}

	return linkHandler.getLink(information, searchResponse), nil
}

// GetAlbum Fetches an album with the apple music identifier
func (linkHandler LinkHandler) GetAlbum(id string) (music.Information, error) {
	musicResponse := music.Information{}
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
func (linkHandler LinkHandler) GetArtist(id string) (music.Information, error) {
	musicResponse := music.Information{}
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
func (linkHandler LinkHandler) GetSong(id string) (music.Information, error) {
	musicResponse := music.Information{}
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

func (linkHandler LinkHandler) mapMediaType(mediaType string) string {
	switch mediaType {
	case "artist":
		return "artists"
	case "album":
		return "albums"
	default:
		return "songs"
	}
}

func (linkHandler LinkHandler) getSearchTerm(information music.Information) string {
	switch information.MediaType {
	case "artist":
		return strings.Replace(information.Artist, " ", "+", -1)
	case "album":
		return strings.Replace(information.Artist+" "+information.Album, " ", "+", -1)
	default:
		return strings.Replace(information.Artist+" "+information.Album+" "+information.Song, " ", "+", -1)
	}
}

func (linkHandler LinkHandler) getLink(response music.Information, searchResponse SearchResponse) string {
	switch response.MediaType {
	case "artist":
		return searchResponse.Results.Artists.Data[0].Attributes.URL
	case "album":
		return searchResponse.Results.Albums.Data[0].Attributes.URL
	default:
		return searchResponse.Results.Songs.Data[0].Attributes.URL
	}
}
