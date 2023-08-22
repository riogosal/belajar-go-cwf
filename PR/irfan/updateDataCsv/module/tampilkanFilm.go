package module

import (
	"belajar-go-cwf/PR/irfan/updateDataCsv/model"
	"fmt"
)

func TampilkanFilm(DataMovies []model.Movie) {
	fmt.Println(len(DataMovies))
	if len(DataMovies) == 0 {

		fmt.Println("Data Film Kosong")
		fmt.Println("")
	}
	fmt.Println("Ini dia koleksi film kamu", len(DataMovies))
	for _, v := range DataMovies {
		v.TampilFilm()
	}

}
