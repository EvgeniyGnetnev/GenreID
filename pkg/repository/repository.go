package repository

type Artist interface {
}

type Genre interface {
}

type Song interface {
}

type Repository struct {
	Artist
	Genre
	Song
}

func NewRepository() *Repository {
	return &Repository{}
}
