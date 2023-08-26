package fungsi

import (
	"belajar-go-cwf/PR/eser/crud/model"
	"fmt"
)

func TampilkanDaftarFilm(koleksi_film map[string]model.Film) {
	fmt.Println("\nDaftar Koleksi Film Kamu :")
	for _, film := range koleksi_film {
		fmt.Println(film.Tampilkan())
	}
}
