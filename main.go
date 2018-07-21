package main

import (
	"context"
	"errors"
	"log"

	"./appleMusic"
	"./music"
	"./spotify"

	"github.com/aws/aws-lambda-go/lambda"
)

type SongBridgeResponse struct {
	Information SongBridgeInformation `json:"information"`
	Links       []SongBridgeLink      `json:"links"`
}

type SongBridgeInformation struct {
	MediaType string `json:"mediaType"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	Song      string `json:"song"`
	Cover     string `json:"cover"`
}

type SongBridgeLink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func getInformations(linkHandler music.LinkHandler, mediaType string, id string) (music.Information, error) {
	switch mediaType {
	case "artist":
		return linkHandler.GetArtist(id)
	case "album":
		return linkHandler.GetAlbum(id)
	default:
		return linkHandler.GetSong(id)
	}
}

type MyEvent struct {
	Source    string `json:"source"`
	MediaType string `json:"mediaType"`
	Id        string `json:"id"`
}

func HandleRequest(ctx context.Context, name MyEvent) (SongBridgeResponse, error) {
	songBrideResponse := SongBridgeResponse{}
	source := name.Source
	if source == "" {
		return songBrideResponse, errors.New("Parameter source is missing")
	}
	mediaType := name.MediaType
	if mediaType == "" {
		return songBrideResponse, errors.New("Parameter mediaType is missing")
	}
	id := name.Id
	if id == "" {
		return songBrideResponse, errors.New("Parameter id is missing")
	}

	log.Print("Source is " + source)
	log.Print("MediaType is " + mediaType)
	log.Print("Id is " + id)

	linkHandler := make(map[string]music.LinkHandler)
	linkHandler["appleMusic"] = applemusic.LinkHandler{}
	linkHandler["spotify"] = spotify.LinkHandler{}

	informations, err := getInformations(linkHandler[source], mediaType, id)
	if err != nil {
		return songBrideResponse, err
	}

	songBrideResponse.Information = SongBridgeInformation{
		MediaType: informations.MediaType,
		Album:     informations.Album,
		Artist:    informations.Artist,
		Song:      informations.Song,
		Cover:     informations.Image,
	}
	for key, value := range linkHandler {
		link, err := value.Search(informations)
		if err != nil {
			continue
		}
		songBrideResponse.Links = append(songBrideResponse.Links, SongBridgeLink{
			Name: key,
			Link: link,
		})
	}

	return songBrideResponse, err
}

func main() {
	lambda.Start(HandleRequest)
}
