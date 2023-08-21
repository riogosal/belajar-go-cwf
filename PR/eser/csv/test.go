package main

import (
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

func main() {
	fileName := "data_film.csv"

	// Buka file CSV
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	// Membaca file CSV sebagai input CSV
	reader := csv.NewReader(file)

	// Inisialisasi slice untuk menyimpan data Film
	var films []Film

	// Membaca baris demi baris dari file CSV
	for {
		record, err := reader.Read()
		if err != nil {
			break // Keluar dari loop jika sudah mencapai akhir file
		}

		// Konversi data dari string ke tipe data yang sesuai
		rating, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Println("Gagal mengonversi rating:", err)
			return
		}

		durasi, err := strconv.Atoi(record[4])
		if err != nil {
			fmt.Println("Gagal mengonversi durasi:", err)
			return
		}

		// Membuat objek Film dari data CSV
		film := Film{
			Judul:     record[0],
			Rating:    rating,
			Deskripsi: record[2],
			Studio:    record[3],
			Durasi:    durasi,
		}

		// Menambahkan objek Film ke dalam slice
		films = append(films, film)
	}

	// Menampilkan data yang sudah dimuat
	fmt.Println("Data Film:")
	for _, film := range films {
		fmt.Printf("Judul: %s\n", film.Judul)
		fmt.Printf("Rating: %f\n", film.Rating)
		fmt.Printf("Deskripsi: %s\n", film.Deskripsi)
		fmt.Printf("Studio: %s\n", film.Studio)
		fmt.Printf("Durasi: %d\n", film.Durasi)
		fmt.Println("-------------")
	}
}
