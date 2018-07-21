package spotify

type AlbumReponse struct {
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Name   string `json:"name"`
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}
