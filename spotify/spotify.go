package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (linkHandler LinkHandler) Search(response music.Response) (string, error) {
	return "", nil
}

// GetArtist Fetches an album with the spotify identifier
func (linkHandler LinkHandler) GetArtist(id string) (music.Response, error) {
	musicResponse := music.Response{}
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
func (linkHandler LinkHandler) GetAlbum(id string) (music.Response, error) {
	musicResponse := music.Response{}
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
func (linkHandler LinkHandler) GetSong(id string) (music.Response, error) {
	musicResponse := music.Response{}
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
