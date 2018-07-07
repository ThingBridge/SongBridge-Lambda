package applemusic

// ArtistResponse Represents apple music artist response
type ArtistResponse struct {
	Data []struct {
		Attributes struct {
			Name string `json:"name"`
		} `json:"attributes"`
	} `json:"data"`
}
