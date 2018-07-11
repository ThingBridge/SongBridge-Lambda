package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// type SongBridgeResponse struct {
// 	Links []SongBridgeLink `json:"links"`
// }

// type SongBridgeLink struct {
// 	Name string `json:"name"`
// 	Link string `json:"link"`
// }

// func getInformations(linkHandler music.LinkHandler, mediaType string, id string) (music.Information, error) {
// 	switch mediaType {
// 	case "artist":
// 		return linkHandler.GetArtist(id)
// 	case "album":
// 		return linkHandler.GetAlbum(id)
// 	default:
// 		return linkHandler.GetSong(id)
// 	}
// }

// func handleBridge(responseWriter http.ResponseWriter, request *http.Request) {
// 	source := request.URL.Query().Get("source")
// 	if source == "" {
// 		responseWriter.WriteHeader(400)
// 		responseWriter.Write([]byte("Query parameter source is missing."))
// 		return
// 	}
// 	mediaType := request.URL.Query().Get("mediaType")
// 	if mediaType == "" {
// 		responseWriter.WriteHeader(400)
// 		responseWriter.Write([]byte("Query parameter mediaType is missing."))
// 		return
// 	}
// 	id := request.URL.Query().Get("id")
// 	if id == "" {
// 		responseWriter.WriteHeader(400)
// 		responseWriter.Write([]byte("Query parameter id is missing."))
// 		return
// 	}

// 	linkHandler := make(map[string]music.LinkHandler)
// 	linkHandler["appleMusic"] = applemusic.LinkHandler{}
// 	linkHandler["spotify"] = spotify.LinkHandler{}

// 	informations, err := getInformations(linkHandler[source], mediaType, id)
// 	if err != nil {
// 		responseWriter.WriteHeader(500)
// 		return
// 	}
// 	songBrideResponse := SongBridgeResponse{}
// 	for key, value := range linkHandler {
// 		link, err := value.Search(informations)
// 		if err != nil {
// 			continue
// 		}
// 		songBrideResponse.Links = append(songBrideResponse.Links, SongBridgeLink{
// 			Name: key,
// 			Link: link,
// 		})
// 	}

// 	data, err := json.Marshal(songBrideResponse)
// 	if err != nil {
// 		responseWriter.WriteHeader(500)
// 		return
// 	}

// 	responseWriter.Header().Set("Content-Type", "application/json")
// 	responseWriter.Write(data)
// }

type MyEvent struct {
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return "Hallo", nil
}

func main() {
	lambda.Start(HandleRequest)
}
