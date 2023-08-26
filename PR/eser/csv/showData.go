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

func updateFilm(film []Film, inputJudul string) {
	for i := range film {
		if film[i].Judul == inputJudul {
			fmt.Printf("Judul saat ini: %s\n", film[i].Judul)
			fmt.Print("Masukkan Judul baru: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newJudul := scanner.Text()
			film[i].Judul = newJudul
			fmt.Println("Judul telah diperbarui.")
			fmt.Println(film)
			return
		}
	}
	fmt.Println("Film tidak ditemukan.")
}

func main() {
	var myFilm []Film
	for {
		fmt.Println("1. Lihat Data")
		fmt.Println("2. Update Data")
		fmt.Println("Pilih menu (1,2), ketik (e) untuk exit")

		fmt.Print("Input Menu :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputMenu := scanner.Text()

		if inputMenu == "e" {
			fmt.Println("Menutup Aplikasi!")
			break
		} else if inputMenu == "1" {
			filename := "data_film.csv"

			file, err := os.Open(filename)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			readData := csv.NewReader(file)

			// kita buat variabel untuk menampung datanya
			myFilm = nil // Mengosongkan slice sebelum mengisi ulang data.

			for {
				index, err := readData.Read()
				if err != nil {
					break // Keluar dari loop jika sudah mencapai akhir file
				}

				// Konversi data dari string ke tipe data yang sesuai
				rating, err := strconv.ParseFloat(index[1], 64)
				if err != nil {
					fmt.Println("Gagal mengonversi rating:", err)
					return
				}

				durasi, err := strconv.Atoi(index[4])
				if err != nil {
					fmt.Println("Gagal mengonversi durasi:", err)
					return
				}

				// Membuat objek Film dari data CSV
				myDataFilm := Film{
					Judul:     index[0],
					Rating:    rating,
					Deskripsi: index[2],
					Studio:    index[3],
					Durasi:    durasi,
				}
				myFilm = append(myFilm, myDataFilm)
			}

			fmt.Println("Data Film:")
			for _, film := range myFilm {
				fmt.Printf("Judul: %s\n", film.Judul)
				fmt.Printf("Rating: %f\n", film.Rating)
				fmt.Printf("Deskripsi: %s\n", film.Deskripsi)
				fmt.Printf("Studio: %s\n", film.Studio)
				fmt.Printf("Durasi: %d\n", film.Durasi)
				fmt.Println("-------------")
			}
		} else if inputMenu == "2" {

			fmt.Print("Masukkan Judul: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			inputJudul := scanner.Text()

			updateFilm(myFilm, inputJudul)
		}
	}
}
