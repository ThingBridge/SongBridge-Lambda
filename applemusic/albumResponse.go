package applemusic

// AlbumResponse Represents a apple music search result
type AlbumResponse struct {
	Data []struct {
		Attributes struct {
			ArtistName string `json:"artistName"`
			Name       string `json:"name"`
		} `json:"attributes"`
	} `json:"data"`
}
