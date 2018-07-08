package music

// LinkHandler Represents a handler for request againt a streaming service
type LinkHandler interface {
	Search(information Information) (string, error)

	GetArtist(id string) (Information, error)
	GetAlbum(id string) (Information, error)
	GetSong(id string) (Information, error)
}
