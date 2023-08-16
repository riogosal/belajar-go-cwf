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
	UUIDStr = strings.TrimSpace(UUIDStr)
	UUID, err := strconv.Atoi(UUIDStr)
	if err != nil {
		fmt.Println("Error:", err)

	}

	data.UUID = UUID

	// Masukkan data Title
	fmt.Print("Masukkan Judul Film favorite: ")
	data.Title, _ = scanner.ReadString('\n')
	data.Title = strings.TrimSpace(data.Title)

	// Masukkan data Deskripsi
	fmt.Print("Masukkan Deskripsi Film favorite: ")
	data.Desc, _ = scanner.ReadString('\n')
	data.Desc = strings.TrimSpace(data.Desc)

	//input Tahun Film Favorit
	fmt.Print("Masukkan Tahun Film Favorite: ")
	YearStr, _ := scanner.ReadString('\n')
	YearStr = strings.TrimSpace(YearStr)
	Year, err := strconv.Atoi(YearStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data.Year = Year

	// Masukkan data Country
	fmt.Print("Masukkan Negara Film favorite: ")
	data.Country, _ = scanner.ReadString('\n')
	data.Country = strings.TrimSpace(data.Country)

	GenreFilms(data)

	fmt.Print("Masukkan Rating Film favorite: ")
	var rate entity.Rating
	fmt.Scanln(&rate)

	data.IMDb = rate

}

func GenreFilms(data *entity.Movies) {
	var no int = 1
	if data.Genre == nil {
		for {
			// Masukkan data Genre
			fmt.Print("Masukkan Genre Film favorite, tekan (e) Exit : ")
			var InputGenre entity.Genre
			fmt.Scanln(&InputGenre)

			if InputGenre == "e" {
				break
			} else {
				data.Genre = append(data.Genre, InputGenre)
			}

		}
	} else {
		for i, v := range data.Genre {
			fmt.Println(no, "Genre : ", v)
			fmt.Println("Update genre : ", no)
			var choiceGenre entity.Genre
			fmt.Scanln(&choiceGenre)
			data.Genre[i] = choiceGenre
			no++
		}
	}
}

func UpdateDataFilmFavorite(data *entity.Movies) {
	scanner := bufio.NewReader(os.Stdin)
myLoops:
	for {
		fmt.Println("Silahakan pilih menu, 1. Ubah Judul, 2. Ubah Deskripsi, 3. Ubah Tahun , 4. Ubah Negara, 5. Ubah Rating, 6. Update Genre, e (exit) :")
		var choiceUpdate string
		fmt.Scanln(&choiceUpdate)
		switch choiceUpdate {
		case "1":
			// Update data Title
			fmt.Print("Update Judul Film favorite: ")
			data.Title, _ = scanner.ReadString('\n')
			data.Title = strings.TrimSpace(data.Title)
		case "2":
			// Update data Deskripsi
			fmt.Print("Update Deskripsi Film favorite: ")
			data.Desc, _ = scanner.ReadString('\n')
			data.Desc = strings.TrimSpace(data.Desc)
		case "3":
			//Update Tahun Film Favorit
			fmt.Print("Update Tahun Film Favorite: ")
			YearStr, _ := scanner.ReadString('\n')
			YearStr = strings.TrimSpace(YearStr)
			Year, err := strconv.Atoi(YearStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			data.Year = Year
		case "4":
			// Update data Country
			fmt.Print("Update Negara Film favorite: ")
			data.Country, _ = scanner.ReadString('\n')
			data.Country = strings.TrimSpace(data.Country)
		case "5":
			fmt.Print("Update Rating Film favorite: ")
			var rate entity.Rating
			fmt.Scanln(&rate)

			data.IMDb = rate
		case "6":
			GenreFilms(data)
		default:
			break myLoops

		}

	}

}

func main() {
	fmt.Println("Selamat datang di program film favorit")
	var TampungDataMovies entity.Movies
	fmt.Println()

myLoop:
	for true {
		fmt.Println("Silahakan pilih menu, 1. tambah data, 2. Update data,3. Hapus Data ,4. Lihat Data :")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if TampungDataMovies.UUID == 0 {
				if TampungDataMovies.UUID == 0 {
					InputDataFilmFavorite(&TampungDataMovies)
				} else {
					fmt.Println("Data Sudah Ada !")
				}

			}
			fmt.Println(" Data Telah Ditambahkan ! ")
		case "2":
			if TampungDataMovies.UUID == 0 {
				fmt.Println(" Data Belum dapat dia Update, silahkan buat Data ! ")
			} else {

				UpdateDataFilmFavorite(&TampungDataMovies)
			}
		case "3":
			TampungDataMovies = entity.Movies{}
			fmt.Println(" Data Telah Dihapus ! ")
		case "4":
			if TampungDataMovies.Title == "" {
				fmt.Println(" Data Film Favorite anda masih kosong, silahkan buat Data ! ")
				TampungDataMovies.PrintFilmFavorite()
			} else {
				TampungDataMovies.PrintFilmFavorite()
			}
		default:
			fmt.Println("Anda Tidak memilih menu apapun !")
			fmt.Println("Arigatou")
			break myLoop
		}
	}

}
