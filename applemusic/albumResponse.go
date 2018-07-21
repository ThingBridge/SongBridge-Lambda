package applemusic

// AlbumResponse Represents a apple music search result
type AlbumResponse struct {
	Data []struct {
		Attributes struct {
			ArtistName string `json:"artistName"`
			Name       string `json:"name"`
			Artwork    struct {
				URL string `json:"url"`
			} `json:"artwork"`
		} `json:"attributes"`
	} `json:"data"`
}
