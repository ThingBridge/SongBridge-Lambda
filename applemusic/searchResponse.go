package applemusic

type SearchResponse struct {
	Results struct {
		Albums struct {
			Data []struct {
				Attributes struct {
					URL string `json:"url"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"albums"`
		Artists struct {
			Data []struct {
				Attributes struct {
					URL string `json:"url"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"artists"`
		Songs struct {
			Data []struct {
				Attributes struct {
					URL string `json:"url"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"songs"`
	} `json:"results"`
}
