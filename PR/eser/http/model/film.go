package model

type FilmDTO struct {
	ID        int     `json:"id"`
	Judul     string  `json:"judul"`
	Deskripsi string  `json:"deskripsi"`
	Studio    string  `json:"studio"`
	Rating    float64 `json:"rating"`
	Durasi    int     `json:"durasi"`
}

type Film struct {
	// STRUCT TAGS
	ID        int     `json:"idFilm"`
	Judul     string  `json:"judul"`
	Deskripsi string  `json:"deskripsi"`
	Studio    string  `json:"studio"`
	Rating    float64 `json:"rating"`
	Durasi    int     `json:"durasi"`
}

var Films []Film = []Film{}
