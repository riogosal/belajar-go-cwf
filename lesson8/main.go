package main

import (
	"belajar-go-cwf/lesson8/cli"
	"belajar-go-cwf/lesson8/fungsi"
	"belajar-go-cwf/lesson8/model"
	"fmt"
)

const DATA_FILE = "films.csv"

func main() {
	koleksi_film := []model.Film{}
	fmt.Println(koleksi_film)

	// 1. Load dulu film
	fungsi.Load(DATA_FILE, &koleksi_film)

	for {
		pilihan_menu := cli.TampilkanMenu()

		if pilihan_menu == cli.TampilkanFilm {
			fungsi.TampilkanFilmFilm(koleksi_film)
		} else if pilihan_menu == cli.BuatFilm {
			fungsi.BuatFilm(&koleksi_film)
		} else if pilihan_menu == cli.UpdateFilm {
			// TODO: MAKE THESE FUNCTIONS
			fmt.Println("Update film")
		} else if pilihan_menu == cli.HapusFilm {
			// TODO: MAKE THESE FUNCTIONS
			fmt.Println("Hapus film")
		} else if pilihan_menu == cli.Exit {
			fungsi.Save(DATA_FILE, koleksi_film)
			break
		} else {
			fmt.Println("Pilihan menu tidak di temukan")
		}
	}

	// TODO: SAVE THE RESULT FROM RAM TO SSD / CSV
}
