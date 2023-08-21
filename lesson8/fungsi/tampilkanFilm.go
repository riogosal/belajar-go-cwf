package fungsi

import (
	"belajar-go-cwf/lesson8/model"
	"fmt"
)

func TampilkanFilmFilm(koleksi_film []model.Film) {
	fmt.Println("Ini dia koleksi film kamu")
	for _, film := range koleksi_film {
		fmt.Println(film.Tampilkan())
	}
}
