package fungsi

import (
	"belajar-go-cwf/PR/eser/crudmap/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func HapusFilm(koleksi_film map[string]model.Film, id_int int) bool {
	for judul, film := range koleksi_film {
		if film.ID == id_int {
			delete(koleksi_film, judul)
			return true
		}
	}
	return false
}

func UserInputId(koleksi_film map[string]model.Film) {
	user_input := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan ID yang ingin Di Hapus : ")
	user_input.Scan()
	id_user_input := user_input.Text()
	id_int, err := strconv.Atoi(id_user_input)
	if err != nil {
		panic(err)
	}

	hapusDataFilm := HapusFilm(koleksi_film, id_int)

	if hapusDataFilm {
		fmt.Println("Film berhasil dihapus!")
	} else {
		fmt.Println("ID Film Tidak di temukan!")
	}
}
