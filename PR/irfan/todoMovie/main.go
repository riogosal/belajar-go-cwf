package main

import (
	"app-movies/entity"
)

func DataHarryPotter() *entity.Movies {
	DataHarryPotter := entity.Movies{
		UUID:    1,
		Title:   "HARRY POTTER AND THE SORCERER’S STONE",
		Desc:    "This is the tale of Harry Potter (Daniel Radcliffe), an ordinary eleven-year-old boy serving as a sort of slave for his aunt and uncle who learns that he is actually a wizard and has been invited to attend the Hogwarts School for Witchcraft and Wizardry.",
		Year:    2001,
		Country: "United States",
		Genre: []entity.Genre{
			{
				Jenis: "Adventure",
			},
			{
				Jenis: "Family",
			},
			{
				Jenis: "Fantasy",
			},
		},
		IMDb: 7.6,
	}

	return &DataHarryPotter
}

func main() {

	film1 := DataHarryPotter()
	film1.CetakMovies()
}
