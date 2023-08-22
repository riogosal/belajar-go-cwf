package main

import (
	"belajar-go-cwf/PR/irfan/updateDataCsv/cli"
	"belajar-go-cwf/PR/irfan/updateDataCsv/model"
	"belajar-go-cwf/PR/irfan/updateDataCsv/module"
	"fmt"
)

const DATA_FILE_NAME = "DataFilm.csv"

func main() {
	DataMovies := []model.Movie{}

	fmt.Println("Data Film Anda : ", DataMovies)

	module.Load(DATA_FILE_NAME, &DataMovies)

	for {
		pilihan_menu := cli.TampilkanMenu()

		switch pilihan_menu {
		case cli.TampilkanFilm:
			module.TampilkanFilm(DataMovies)
		case cli.BuatFilm:
			module.BuatFilm(&DataMovies)
		case cli.UpdateFilm:
			module.UpdateFilm(DataMovies)
		case cli.HapusFilm:
			DataMovies = []model.Movie{}
		case cli.Exit:
			module.Save(DATA_FILE_NAME, DataMovies)
			fmt.Println("Data Telah Succes Diimport")
			break
		default:
			fmt.Println("anda tidak memilih menu apapun")
			break

		}
	}
}
