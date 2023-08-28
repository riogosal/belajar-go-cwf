package fungsi

import (
	"belajar-go-cwf/PR/eser/crud/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TanyaUserFilmBaru() model.Film {
	user_input := bufio.NewScanner(os.Stdin)

	fmt.Print("ID	    : ")
	user_input.Scan()
	id_user_input := user_input.Text()
	id_int, err := strconv.Atoi(id_user_input)
	if err != nil {
		panic(err)
	}

	fmt.Print("Judul     : ")
	user_input.Scan()
	judul := user_input.Text()

	fmt.Print("Deskripsi : ")
	user_input.Scan()
	deskripsi_user_input := user_input.Text()

	fmt.Print("Studio    : ")
	user_input.Scan()
	studio_user_input := user_input.Text()

	fmt.Print("Rating    : ")
	user_input.Scan()
	rating_user_input := user_input.Text()
	rating_float, err := strconv.ParseFloat(rating_user_input, 64)
	if err != nil {
		panic(err)
	}

	fmt.Print("Durasi    : ")
	user_input.Scan()
	durasi_user_input := user_input.Text()
	durasi_int, err := strconv.Atoi(durasi_user_input)
	if err != nil {
		panic(err)
	}

	return model.Film{
		ID:        id_int,
		Judul:     judul,
		Deskripsi: deskripsi_user_input,
		Studio:    studio_user_input,
		Rating:    rating_float,
		Durasi:    durasi_int,
	}
}

func BuatFilm(koleksi_film *[]model.Film) {
	*koleksi_film = append(*koleksi_film, TanyaUserFilmBaru())
}
