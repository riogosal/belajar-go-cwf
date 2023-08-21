package fungsi

import (
	"belajar-go-cwf/lesson8/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TanyaUserFilmBaru() model.Film {
	user_input := bufio.NewScanner(os.Stdin)

	fmt.Print("Judul: ")
	user_input.Scan()
	judul := user_input.Text()

	fmt.Print("Deskripsi: ")
	user_input.Scan()
	deskripsi_dari_user_input := user_input.Text()

	fmt.Print("Durasi: ")
	user_input.Scan()
	durasi_dari_user_input := user_input.Text()
	durasi_int, err := strconv.Atoi(durasi_dari_user_input)
	if err != nil {
		panic(err)
	}

	return model.Film{
		Judul:     judul,
		Deskripsi: deskripsi_dari_user_input,
		Durasi:    durasi_int,
	}
}

func BuatFilm(koleksi_film *[]model.Film) {
	*koleksi_film = append(*koleksi_film, TanyaUserFilmBaru())
}
