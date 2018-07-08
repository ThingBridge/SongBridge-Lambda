package spotify

type TrackResponse struct {
	Album struct {
		Name string `json:"name"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Name string `json:"name"`
}
