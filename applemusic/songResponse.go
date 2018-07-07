package applemusic

// SongResponse Repesents the response from an apple music song request
type SongResponse struct {
	Data []struct {
		Attributes struct {
			ArtistName string `json:"artistName"`
			AlbumName  string `json:"albumName"`
			Name       string `json:"name"`
		} `json:"attributes"`
	} `json:"data"`
}
