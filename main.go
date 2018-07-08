package main

import (
	"flag"
	"fmt"

	"./applemusic"
	"./music"
	"./spotify"
)

func main() {
	source := flag.String("source", "", "Sets the source of the id")
	id := flag.String("id", "", "The id too lookup in other sevices")
	mediaType := flag.String("mediaType", "", "Kind of media")
	flag.Parse()

	var linkHandler music.LinkHandler = applemusic.LinkHandler{}
	var targetLinkHandler music.LinkHandler = spotify.LinkHandler{}
	if *source == "spotify" {
		linkHandler = spotify.LinkHandler{}
		targetLinkHandler = applemusic.LinkHandler{}
	}

	if *mediaType == "artist" {
		response, err := linkHandler.GetArtist(*id)
		if err != nil {
			panic(err)
		}
		link, err := targetLinkHandler.Search(response)
		if err != nil {
			panic(err)
		}
		fmt.Println(link)
	} else if *mediaType == "album" {
		response, err := linkHandler.GetAlbum(*id)
		if err != nil {
			panic(err)
		}
		link, err := targetLinkHandler.Search(response)
		if err != nil {
			panic(err)
		}
		fmt.Println(link)
	} else {
		response, err := linkHandler.GetSong(*id)
		if err != nil {
			panic(err)
		}
		link, err := targetLinkHandler.Search(response)
		if err != nil {
			panic(err)
		}
		fmt.Println(link)
	}

	// // Spotify
	// response, err := spotifyMusicLinkHandler.GetArtist("3AA28KZvwAUcZuOKwyblJQ")
	// if err != nil {
	// 	panic(err)
	// }

	// link, err := appleMusicLinkHandler.Search(response)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(link)

	// response, err = spotifyMusicLinkHandler.GetAlbum("1amYhlukNF8WdaQC3gKkgL")
	// if err != nil {
	// 	panic(err)
	// }

	// link, err = appleMusicLinkHandler.Search(response)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(link)

	// response, err = spotifyMusicLinkHandler.GetSong("7jYUaoOfdcYgUvkK8NnFfx")
	// if err != nil {
	// 	panic(err)
	// }

	// link, err = appleMusicLinkHandler.Search(response)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(link)

	// // Apple Music
	// response, err = appleMusicLinkHandler.GetArtist("567072")
	// if err != nil {
	// 	panic(err)
	// }
	// link, err = spotifyMusicLinkHandler.Search(response)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(link)

	// response, err = appleMusicLinkHandler.GetAlbum("1387814084")
	// if err != nil {
	// 	panic(err)
	// }
	// link, err = spotifyMusicLinkHandler.Search(response)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(link)

	// response, err = appleMusicLinkHandler.GetSong("1387814980")
	// if err != nil {
	// 	panic(err)
	// }

	// link, err = spotifyMusicLinkHandler.Search(response)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(link)
}
