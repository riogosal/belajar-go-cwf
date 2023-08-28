package fungsi

import (
	"belajar-go-cwf/PR/eser/crud/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpdateFilm(koleksi_film *[]model.Film) {
	TampilkanDaftarFilm(*koleksi_film)
	// fmt.Println(koleksi_film)

	ListFilmToMap := make(map[string]model.Film)

	for _, film := range *koleksi_film {
		ListFilmToMap[film.Judul] = film
	}

	user_input_text := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Judul Yang Ingin Di update : ")
	user_input_text.Scan()
	judulFilm := user_input_text.Text()

	film, found := ListFilmToMap[judulFilm]

	if found {
		fmt.Println("\nJudul    :", film.Judul)
		fmt.Println("Deskripsi:", film.Deskripsi)
		fmt.Println("Studio   :", film.Studio)
		fmt.Println("Rating   :", film.Rating)
		fmt.Println("Durasi   :", film.Durasi)

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("Silahkan pilih menu, 1. Update Judul, 2. Update Rating 3.Update Deskripsi 4. Update Studio, 5.Update Durasi, (e) exit")
			fmt.Print("Pilih Menu :")
			userInput, _ := reader.ReadString('\n')
			userInput = strings.TrimSpace(userInput)

			switch userInput {
			case "1":
				fmt.Print("Masukkan Judul Baru: ")
				judulBaru, _ := reader.ReadString('\n')
				judulBaru = strings.TrimSpace(judulBaru)

				for i, film := range *koleksi_film {
					if film.Judul == judulFilm {
						(*koleksi_film)[i].Judul = judulBaru
						fmt.Println("Judul berhasil diubah:", judulBaru)
						break // Keluar dari loop setelah judul ditemukan dan diubah
					}
				}
				return

			case "2":
				fmt.Print("Masukkan Deskripsi : ")
				deskripsiBaru, _ := reader.ReadString('\n')
				deskripsiBaru = strings.TrimSpace(deskripsiBaru)

				for i, film := range *koleksi_film {
					if film.Deskripsi == deskripsiBaru {
						(*koleksi_film)[i].Deskripsi = deskripsiBaru
						fmt.Println("Deskripsi berhasil diubah:", deskripsiBaru)
						break
					}
				}
				return

			case "3":
				fmt.Print("Masukkan Studio : ")
				studiobaru, _ := reader.ReadString('\n')
				studiobaru = strings.TrimSpace(studiobaru)

				for i, film := range *koleksi_film {
					if film.Deskripsi == studiobaru {
						(*koleksi_film)[i].Studio = studiobaru
						fmt.Println("Studio berhasil diubah:", studiobaru)
						break
					}
				}
				return
			case "4":
				fmt.Print("Masukkan Rating: ")
				ratingBaruStr, _ := reader.ReadString('\n')
				ratingBaruStr = strings.TrimSpace(ratingBaruStr)
				ratingFloat, err := strconv.ParseFloat(ratingBaruStr, 64)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				for i, film := range *koleksi_film {
					if film.Rating == ratingFloat {
						(*koleksi_film)[i].Rating = ratingFloat
						fmt.Println("Rating berhasil diubah:", ratingFloat)
						break
					}
				}
				return

			case "e":
				return
			}
			fmt.Println("Data Film Setelah Perubahan Judul:")
			for _, film := range *koleksi_film {
				fmt.Println("ID	      : ", film.ID)
				fmt.Println("Judul    : ", film.Judul)
				fmt.Println("Deskripsi: ", film.Deskripsi)
				fmt.Println("Studio   : ", film.Studio)
				fmt.Println("Rating   : ", film.Rating)
				fmt.Println("Durasi   : ", film.Durasi)
				fmt.Println("------------------")
			}
		}
	} else {
		fmt.Println("Film dengan judul", judulFilm, "tidak ditemukan.")
	}

}
