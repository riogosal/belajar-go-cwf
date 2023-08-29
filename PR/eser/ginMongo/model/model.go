package model

type Genre string

type Film struct {
	// ID        int     `bson:"id" validate:"required,unique"`
	ID        int     `bson:"id"`
	Judul     string  `bson:"judul"`
	Deskripsi string  `bson:"deskripsi"`
	Studio    string  `bson:"studio"`
	Rating    float64 `bson:"rating"`
	Durasi    int     `bson:"durasi"`
	Genres    []Genre `bson:"genres"`
}

// var Films []Film = []Film{}
var Films []Film
