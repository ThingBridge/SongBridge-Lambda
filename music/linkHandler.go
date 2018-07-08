package music

// LinkHandler Represents a handler for request againt a streaming service
type LinkHandler interface {
	Search(response Response) (string, error)

	GetArtist(id string) (Response, error)
	GetAlbum(id string) (Response, error)
	GetSong(id string) (Response, error)
}
