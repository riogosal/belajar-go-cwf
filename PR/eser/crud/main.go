package main

import (
	"belajar-go-cwf/PR/eser/crud/cli"
	"belajar-go-cwf/PR/eser/crud/fungsi"
	"belajar-go-cwf/PR/eser/crud/model"
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
			fungsi.TampilkanDaftarFilm(koleksi_film)
		} else if pilihan_menu == cli.BuatFilm {
			fungsi.BuatFilm(&koleksi_film)
		} else if pilihan_menu == cli.UpdateFilm {
			fungsi.UpdateFilm(&koleksi_film)
		} else if pilihan_menu == cli.HapusFilm {
			// fmt.Println("Hapus film")
			fungsi.HapusDataFilm(&koleksi_film)
		} else if pilihan_menu == cli.Exit {
			fungsi.Save(DATA_FILE, koleksi_film)
			break
		} else {
			fmt.Println("Pilihan menu tidak di temukan")
		}
	}

	// TODO: SAVE THE RESULT FROM RAM TO SSD / CSV
}
