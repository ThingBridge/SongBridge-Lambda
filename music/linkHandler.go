package music

// LinkHandler Represents a handler for request againt a streaming service
type LinkHandler interface {
	GetArtist(id string) (Response, error)
	GetAlbum(id string) (Response, error)
	GetSong(id string) (Response, error)
}
