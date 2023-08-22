package fungsi

import (
	"belajar-go-cwf/PR/eser/crudmap/model"
	"fmt"
)

func TampilkanDataFilm(koleksi_film map[string]model.Film) {
	fmt.Println(koleksi_film)
	fmt.Println("\nDaftar Koleksi Film Kamu :")
	for _, film := range koleksi_film {
		fmt.Println(film.Tampilkan())
	}
}
