package genreid

type Artist struct {
	Id int `json:"-"`
	Name string `json:"name"`
}

type Genre struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desription"`
}

type Song struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desription"`
}

type AtristSong struct {
	Id int
	ArtistId int
	SongId int
}

type SongGenre struct {
	Id int
	SongtId int
	GenreId int
}