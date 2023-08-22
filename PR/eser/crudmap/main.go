package main

import (
	// "belajar-go-cwf/PR/eser/crudmap/cli"
	"belajar-go-cwf/PR/eser/crudmap/cli"
	"belajar-go-cwf/PR/eser/crudmap/fungsi"
	"belajar-go-cwf/PR/eser/crudmap/model"
	"fmt"
)

const DATA_FILE = "films.csv"

func main() {
	koleksi_film := make(map[string]model.Film)

	// Load data dari file CSV dan masukkan ke dalam map
	fungsi.Load(DATA_FILE, koleksi_film)

	for {
		pilihan_menu := cli.TampilkanMenu()

		if pilihan_menu == cli.TampilkanFilm {
			fungsi.TampilkanDataFilm(koleksi_film)
		} else if pilihan_menu == cli.BuatFilm {
			fungsi.BuatFilm(koleksi_film)
		} else if pilihan_menu == cli.Exit {
			fungsi.Save(DATA_FILE, koleksi_film)
			break
		} else if pilihan_menu == cli.UpdateFilm {
			fungsi.UpdateFilm(koleksi_film)
		} else if pilihan_menu == cli.HapusFilm {
			// fmt.Println("Hapus film")
			fungsi.UserInputId(&koleksi_film)
		} else {
			fmt.Println("Pilihan menu tidak di temukan")
		}
	}

	// TODO: SAVE THE RESULT FROM RAM TO SSD / CSV
}
