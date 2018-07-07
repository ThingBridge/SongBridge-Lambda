package main

import (
	"fmt"
	"log"

	"./applemusic"
	"./music"
)

func main() {
	linkHandler := make(map[string]music.LinkHandler)
	linkHandler["applemusic"] = applemusic.LinkHandler{}

	response, err := linkHandler["applemusic"].GetAlbum("617154241")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response.MediaType)
		fmt.Println(response.Artist)
		fmt.Println(response.Album)
		fmt.Println(response.Song)
	}

	response, err = linkHandler["applemusic"].GetArtist("5468295")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response.MediaType)
		fmt.Println(response.Artist)
		fmt.Println(response.Album)
		fmt.Println(response.Song)
	}

	response, err = linkHandler["applemusic"].GetSong("1156443304")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response.MediaType)
		fmt.Println(response.Artist)
		fmt.Println(response.Album)
		fmt.Println(response.Song)
	}
}
