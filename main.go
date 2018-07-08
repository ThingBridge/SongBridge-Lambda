package main

import (
	"fmt"

	"./applemusic"
	"./spotify"
)

func main() {
	appleMusicLinkHandler := applemusic.LinkHandler{}
	spotifyMusicLinkHandler := spotify.LinkHandler{}

	// Spotify
	response, err := spotifyMusicLinkHandler.GetArtist("3AA28KZvwAUcZuOKwyblJQ")
	if err != nil {
		panic(err)
	}

	link, err := appleMusicLinkHandler.Search(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(link)

	response, err = spotifyMusicLinkHandler.GetAlbum("1amYhlukNF8WdaQC3gKkgL")
	if err != nil {
		panic(err)
	}

	link, err = appleMusicLinkHandler.Search(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(link)

	response, err = spotifyMusicLinkHandler.GetSong("7jYUaoOfdcYgUvkK8NnFfx")
	if err != nil {
		panic(err)
	}

	link, err = appleMusicLinkHandler.Search(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(link)

	// Apple Music
	response, err = appleMusicLinkHandler.GetArtist("567072")
	if err != nil {
		panic(err)
	}
	link, err = spotifyMusicLinkHandler.Search(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(link)

	response, err = appleMusicLinkHandler.GetAlbum("1387814084")
	if err != nil {
		panic(err)
	}
	link, err = spotifyMusicLinkHandler.Search(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(link)

	response, err = appleMusicLinkHandler.GetSong("1387814980")
	if err != nil {
		panic(err)
	}

	link, err = spotifyMusicLinkHandler.Search(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(link)
}
