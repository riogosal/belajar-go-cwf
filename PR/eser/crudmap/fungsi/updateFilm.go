package fungsi

import (
	"belajar-go-cwf/PR/eser/crudmap/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UpdateFilm(koleksi_film map[string]model.Film) {
	fmt.Println(koleksi_film)

	fmt.Println("Data Film favorit kamu :")
	for judul, film := range koleksi_film {
		fmt.Println("Judul    :", judul)
		fmt.Println("ID       :", film.ID)
		fmt.Println("Deskripsi:", film.Deskripsi)
		fmt.Println("Studio   :", film.Studio)
		fmt.Println("Rating   :", film.Rating)
		fmt.Println("Durasi   :", film.Durasi)
		fmt.Println("------------------")
	}

	user_input_text := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Judul Yang Ingin Diupdate: ")
	user_input_text.Scan()
	judulFilm := user_input_text.Text()

	film, found := koleksi_film[judulFilm]

	if found {
		fmt.Println("\nFilm Berhasil Ditemukan!")
		fmt.Println("Judul    :", film.Judul)
		fmt.Println("Deskripsi:", film.Deskripsi)
		fmt.Println("Studio   :", film.Studio)
		fmt.Println("Rating   :", film.Rating)
		fmt.Println("Durasi   :", film.Durasi)

		for {
			fmt.Println("\nSilahkan pilih menu:")
			fmt.Println("1. Update Judul")
			fmt.Println("2. Update Deskripsi")
			fmt.Println("3. Update Studio")
			fmt.Println("4. Update Rating")

			fmt.Print("Pilih Menu: ")
			userInputText := bufio.NewScanner(os.Stdin)
			userInputText.Scan()
			userInput := userInputText.Text()

			switch userInput {
			case "1":
				fmt.Print("Masukkan Judul Baru: ")
				userInputText.Scan()
				judulBaru := userInputText.Text()
				film.Judul = judulBaru
				koleksi_film[judulFilm] = film
				fmt.Println("Judul berhasil diubah:", judulBaru)
				return
			case "2":
				fmt.Print("Masukkan Deskripsi Baru: ")
				userInputText.Scan()
				deskripsiBaru := userInputText.Text()
				film.Deskripsi = deskripsiBaru
				koleksi_film[judulFilm] = film
				fmt.Println("Deskripsi berhasil diubah:", deskripsiBaru)
				return
			case "3":
				fmt.Print("Masukkan Studio Baru: ")
				userInputText.Scan()
				studioBaru := userInputText.Text()
				film.Studio = studioBaru
				koleksi_film[judulFilm] = film
				fmt.Println("Studio berhasil diubah:", studioBaru)
				return
			case "4":
				fmt.Print("Masukkan Rating Baru: ")
				userInputText.Scan()
				ratingBaruText := userInputText.Text()
				ratingBaru, err := strconv.ParseFloat(ratingBaruText, 64)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				film.Rating = ratingBaru
				koleksi_film[judulFilm] = film
				fmt.Println("Rating berhasil diubah:", ratingBaru)
				return
			default:
				fmt.Println("Menu yang di pilih tidak ada")
			}
		}
	} else {
		fmt.Println("Judul Film tidak ditemukan!")
	}
}
