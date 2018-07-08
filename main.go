package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./applemusic"
	"./music"
	"./spotify"
)

type SongBridgeResponse struct {
	Links []SongBridgeLink `json:"links"`
}

type SongBridgeLink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func handleBridge(responseWriter http.ResponseWriter, request *http.Request) {
	source := request.URL.Query().Get("source")
	if source == "" {
		responseWriter.WriteHeader(400)
		responseWriter.Write([]byte("Query parameter source is missing."))
		return
	}
	mediaType := request.URL.Query().Get("mediaType")
	if mediaType == "" {
		responseWriter.WriteHeader(400)
		responseWriter.Write([]byte("Query parameter mediaType is missing."))
		return
	}
	id := request.URL.Query().Get("id")
	if id == "" {
		responseWriter.WriteHeader(400)
		responseWriter.Write([]byte("Query parameter id is missing."))
		return
	}

	var linkHandler music.LinkHandler = applemusic.LinkHandler{}
	var targetLinkHandler music.LinkHandler = spotify.LinkHandler{}
	if source == "spotify" {
		linkHandler = spotify.LinkHandler{}
		targetLinkHandler = applemusic.LinkHandler{}
	}

	if mediaType == "artist" {
		response, err := linkHandler.GetArtist(id)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}
		link, err := targetLinkHandler.Search(response)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}
		songBrideResponse := SongBridgeResponse{
			Links: []SongBridgeLink{
				SongBridgeLink{
					Name: "Spotify",
					Link: link,
				},
			},
		}
		data, err := json.Marshal(songBrideResponse)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(data)
	} else if mediaType == "album" {
		response, err := linkHandler.GetAlbum(id)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}
		link, err := targetLinkHandler.Search(response)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}
		songBrideResponse := SongBridgeResponse{
			Links: []SongBridgeLink{
				SongBridgeLink{
					Name: "Spotify",
					Link: link,
				},
			},
		}
		data, err := json.Marshal(songBrideResponse)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(data)
	} else {
		response, err := linkHandler.GetSong(id)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}
		link, err := targetLinkHandler.Search(response)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}
		songBrideResponse := SongBridgeResponse{
			Links: []SongBridgeLink{
				SongBridgeLink{
					Name: "Spotify",
					Link: link,
				},
			},
		}
		data, err := json.Marshal(songBrideResponse)
		if err != nil {
			responseWriter.WriteHeader(500)
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(data)
	}

}

func main() {
	http.HandleFunc("/bridge", handleBridge)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// source := flag.String("source", "", "Sets the source of the id")
	// id := flag.String("id", "", "The id too lookup in other sevices")
	// mediaType := flag.String("mediaType", "", "Kind of media")
	// flag.Parse()

	// var linkHandler music.LinkHandler = applemusic.LinkHandler{}
	// var targetLinkHandler music.LinkHandler = spotify.LinkHandler{}
	// if *source == "spotify" {
	// 	linkHandler = spotify.LinkHandler{}
	// 	targetLinkHandler = applemusic.LinkHandler{}
	// }

	// if *mediaType == "artist" {
	// 	response, err := linkHandler.GetArtist(*id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	link, err := targetLinkHandler.Search(response)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(link)
	// } else if *mediaType == "album" {
	// 	response, err := linkHandler.GetAlbum(*id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	link, err := targetLinkHandler.Search(response)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(link)
	// } else {
	// 	response, err := linkHandler.GetSong(*id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	link, err := targetLinkHandler.Search(response)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(link)
	// }
}
