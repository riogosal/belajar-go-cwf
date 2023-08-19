package main

import (
	"belajar-go-cwf/lesson7/models"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file for reading
	file, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // will execute before the function main() exits.

	fmt.Print("Masukkan judul yang kamu cari: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	judulYangDiCari := scanner.Text()

	// Create a new CSV reader reading from the opened file
	reader := bufio.NewScanner(file)

	koleksiFilm := []models.Film{}
	for reader.Scan() {
		line := reader.Text()

		koleksiFilm = append(
			koleksiFilm,
			models.FromCSV(line), // convert string to Film struct
		)
	}

	fmt.Println("Akhirnya isi koleksi film", koleksiFilm)

	for _, film := range koleksiFilm {
		if film.Judul == judulYangDiCari {
			fmt.Println("Film yang kamu cari di temukan ini datanya", film)
		}
	}
}
