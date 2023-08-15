package main

import (
	"app-data-movies/entity"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputDataFilmFavorite(data *entity.Movies) {
	scanner := bufio.NewReader(os.Stdin)

	//input ID Fil Favorit
	fmt.Print("Masukkan ID Film Favorite: ")
	UUIDStr, _ := scanner.ReadString('\n')
	UUID, err := strconv.Atoi(UUIDStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data.UUID = UUID

	// Masukkan data Title
	fmt.Print("Masukkan ID Film favorite: ")
	data.Title, _ = scanner.ReadString('\n')
	data.Title = strings.TrimSpace(data.Title)

}
func UpdateDataFilmFavorite(data *entity.Movies) {
	data.Title = "Superman"

}

func main() {
	fmt.Println("Selamat datang di program film favorit")
	var TampungDataMovies entity.Movies

myLoop:
	for true {
		fmt.Println("SIlahakan pilih menu, 1. tambah data, 2. Update data,3. Hapus Data ,4. Lihat Data :")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			InputDataFilmFavorite(&TampungDataMovies)
		case "2":
			UpdateDataFilmFavorite(&TampungDataMovies)
		case "3":
			TampungDataMovies = entity.Movies{}
		case "4":
			TampungDataMovies.PrintFilmFavorite()
		default:
			fmt.Println("Anda Tidak memilih menu apapun !")
			fmt.Println("Arigatou")
			break myLoop
		}
	}

}
