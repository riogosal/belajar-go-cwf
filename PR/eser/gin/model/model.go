package model

type Genre string

type Film struct {
	ID        int     `json:"id"`
	Judul     string  `json:"judul"`
	Deskripsi string  `json:"deskripsi"`
	Studio    string  `json:"studio"`
	Rating    float64 `json:"rating"`
	Durasi    int     `json:"durasi"`
	Genres    []Genre `json:"genres"`
}

var Films []Film = []Film{}
