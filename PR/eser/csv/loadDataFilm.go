package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Film struct {
	Judul     string
	Rating    float64
	Deskripsi string
	Studio    string
	Durasi    int
}

func showDataFilm(film *Film) {
	if film.Judul != "" {
		fmt.Println("\nDaftar Film yang kamu cari : ")
		fmt.Println("Judul     :", film.Judul)
		fmt.Println("Deskripsi :", film.Deskripsi)
		fmt.Println("Studio    :", film.Studio)
		fmt.Println("Durasi    :", film.Durasi)
	} else {
		fmt.Println("Error, Anda Salah Koding!!")
	}
}

func main() {

	var dataFilm []Film
	for {
		fmt.Println("\nPilih Menu, (e) Exit: ")
		fmt.Println("1. Cari Judul")
		fmt.Println("2. Lihat Film")
		fmt.Println("3. Update Film")
		fmt.Print("pilih Menu :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputDataScan := scanner.Text()

		if inputDataScan == "e" {
			fmt.Println("Menutup aplikasi")
			break
		} else if inputDataScan == "1" {
			fmt.Print("Masukkan Judul  Film : ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			inputJudulFilm := scanner.Text()

			fileName := "data_film.csv"
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Gagal membuka file:", err)
				return
			}
			defer file.Close()
			/** ref
			https://www.educative.io/answers/how-to-read-a-csv-file-in-the-go-language
			https://www.geeksforgeeks.org/how-to-read-a-csv-file-in-golang/
			*/
			reader := csv.NewReader(file)

			var myFilm Film

			// Membaca file line per line dari file CSV
			for {
				dataFilm, err := reader.Read()
				if err != nil {
					break
				}

				judul := dataFilm[0]

				// jika judulnya sama dengan data yang di input di atas tadi maka ambil semua datanya
				if judul == inputJudulFilm {
					rating, _ := strconv.ParseFloat(dataFilm[1], 64)
					durasi, _ := strconv.Atoi(dataFilm[4])

					myFilm = Film{
						Judul:     judul,
						Rating:    rating,
						Deskripsi: dataFilm[2],
						Studio:    dataFilm[3],
						Durasi:    durasi,
					}

					fmt.Printf("\nJudul    : %s\n", judul)
					fmt.Printf("Rating   : %f\n", rating)
					fmt.Printf("Deskripsi: %s\n", dataFilm[2])
					fmt.Printf("Studio   : %s\n", dataFilm[3])
					fmt.Printf("Durasi   : %d\n", durasi)
					break
				}
			}

			if myFilm.Judul != "" {
				dataFilm = append(dataFilm, myFilm) // Simpan film yang ditemukan dalam slice
			} else {
				fmt.Println("Film tidak ditemukan!")
			}
		} else if inputDataScan == "2" {
			if len(dataFilm) > 0 {
				showDataFilm(&dataFilm[0])
			} else {
				fmt.Println("Belum ada data film yang ditemukan.")
			}
		} else if inputDataScan == "3" {

		}

	}
}
