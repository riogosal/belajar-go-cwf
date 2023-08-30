package module

import (
	"belajar-go-cwf/PR/irfan/updateDataCsv/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InputDataFilm() model.Movie {
	films_input := bufio.NewScanner(os.Stdin)

	fmt.Print("ID: ")
	films_input.Scan()
	data_film_id := films_input.Text()
	id_int, err := strconv.Atoi(data_film_id)
	if err != nil {
		panic(err)
	}

	fmt.Print("Judul: ")
	films_input.Scan()
	judul := films_input.Text()

	fmt.Print("Deskripsi: ")
	films_input.Scan()
	desc := films_input.Text()

	fmt.Print("Tahun: ") 
	films_input.Scan()
	data_film_tahun := films_input.Text()
	tahun_int, err := strconv.Atoi(data_film_tahun)
	if err != nil {
		panic(err)
	}

	return model.Movie{
		Id:        id_int,
		Judul:     judul,
		Deskripsi: desc,
		Tahun:     tahun_int,
	}
}

func BuatFilm(dataMovie *[]model.Movie) {
	*dataMovie = append(*dataMovie, InputDataFilm())
}
