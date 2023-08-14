package main

import "fmt"

type Genre string

type Film struct {
	Genres []Genre
}

func (g Genre) JelaskanGenre() {
	fmt.Println("Genre ini adalah", g)
}

func main() {
	genreHorror := Genre("Horror")
	genreAction := Genre("Action")

	filmSeramDanAction := Film{
		Genres: []Genre{
			genreHorror,
			genreAction,
			"Comedy", // ini bisa juga langsung di isi string tapi tipenya bukan string tapi Genre
		},
	}

	fmt.Println(filmSeramDanAction)

	for _, genre := range filmSeramDanAction.Genres {
		genre.JelaskanGenre()
	}

}
