package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"../music"
	"github.com/go-redis/redis"
)

func getAccessToken() (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	accessToken, err := client.Get("spotify_access_token").Result()
	if err != nil {
		return "", nil
	}
	return accessToken, nil
}

// LinkHandler Represents a handler for requests against apple music
type LinkHandler struct {
}

func (linkHandler LinkHandler) Search(information music.Information) (string, error) {
	types := linkHandler.mapMediaType(information.MediaType)
	searchTerm := linkHandler.getSearchTerm(information)

	accessToken, err := getAccessToken()
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?limit=1&q="+searchTerm+"&type="+types, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

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

// GetArtist Fetches an album with the spotify identifier
func (linkHandler LinkHandler) GetArtist(id string) (music.Information, error) {
	musicResponse := music.Information{}
	spotifyReponse := ArtistResponse{}

	accessToken, err := getAccessToken()
	if err != nil {
		return musicResponse, err
	}

	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/artists/"+id, nil)
	if err != nil {
		return musicResponse, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

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

	err = json.Unmarshal(data, &spotifyReponse)
	if err != nil {
		return musicResponse, err
	}

	musicResponse.MediaType = "artist"
	musicResponse.Artist = spotifyReponse.Name

	return musicResponse, nil
}

// GetAlbum Fetches an album with the spotify identifier
func (linkHandler LinkHandler) GetAlbum(id string) (music.Information, error) {
	musicResponse := music.Information{}
	spotifyReponse := AlbumReponse{}

	accessToken, err := getAccessToken()
	if err != nil {
		return musicResponse, err
	}

	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/albums/"+id, nil)
	if err != nil {
		return musicResponse, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

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

	err = json.Unmarshal(data, &spotifyReponse)
	if err != nil {
		return musicResponse, err
	}

	musicResponse.MediaType = "album"
	musicResponse.Artist = spotifyReponse.Artists[0].Name
	musicResponse.Album = spotifyReponse.Name

	return musicResponse, nil
}

// GetSong Fetches informations for a spotify song
func (linkHandler LinkHandler) GetSong(id string) (music.Information, error) {
	musicResponse := music.Information{}
	spotifyReponse := TrackResponse{}

	accessToken, err := getAccessToken()
	if err != nil {
		return musicResponse, err
	}

	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/tracks/"+id, nil)
	if err != nil {
		return musicResponse, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

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

	err = json.Unmarshal(data, &spotifyReponse)
	if err != nil {
		return musicResponse, err
	}

	musicResponse.MediaType = "song"
	musicResponse.Artist = spotifyReponse.Artists[0].Name
	musicResponse.Album = spotifyReponse.Album.Name
	musicResponse.Song = spotifyReponse.Name

	return musicResponse, nil
}

func (linkHandler LinkHandler) mapMediaType(mediaType string) string {
	switch mediaType {
	case "artist":
		return "artist"
	case "album":
		return "album"
	default:
		return "track"
	}
}

func (linkHandler LinkHandler) getSearchTerm(response music.Information) string {
	switch response.MediaType {
	case "artist":
		return strings.Replace(response.Artist, " ", "+", -1)
	case "album":
		return strings.Replace(response.Artist+" "+response.Album, " ", "+", -1)
	default:
		return strings.Replace(response.Artist+" "+response.Album+" "+response.Song, " ", "+", -1)
	}
}

func (linkHandler LinkHandler) getLink(response music.Information, searchResponse SearchResponse) string {
	switch response.MediaType {
	case "artist":
		return searchResponse.Artists.Items[0].ExternalUrls.Spotify
	case "album":
		return searchResponse.Albums.Items[0].ExternalUrls.Spotify
	default:
		return searchResponse.Tracks.Items[0].ExternalUrls.Spotify
	}
}
