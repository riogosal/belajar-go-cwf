package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Film struct {
	Judul             string
	Rating            float64
	Deskripsi, Studio string
	Durasi            int
	Gendre            string
}

func InputData1(film *Film) {
	// var newData Film
	newData := Film{}

	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Judul: ")
	newData.Judul, _ = scanner.ReadString('\n')
	newData.Judul = strings.TrimSpace(newData.Judul)

	fmt.Print("Masukkan Rating: ")
	ratingStr, _ := scanner.ReadString('\n')
	ratingStr = strings.TrimSpace(ratingStr) // menghilangkan karakter '\n'
	newRating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newData.Rating = newRating

	fmt.Print("Masukkan Deskripsi : ")
	newData.Deskripsi, _ = scanner.ReadString('\n')
	newData.Deskripsi = strings.TrimSpace(newData.Deskripsi)

	fmt.Print("Masukkan Studio : ")
	newData.Studio, _ = scanner.ReadString('\n')
	newData.Studio = strings.TrimSpace(newData.Studio)

	fmt.Print("Masukkan durasi : ")
	durasiStr, _ := scanner.ReadString('\n')
	durasiStr = strings.TrimSpace(durasiStr)
	newDurasi, err := strconv.Atoi(durasiStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newData.Durasi = newDurasi

	fmt.Print("Masukkan Gendre : ")
	newData.Gendre, _ = scanner.ReadString('\n')
	newData.Gendre = strings.TrimSpace(newData.Gendre)

	// untuk akses value dari pointer
	// musti di dereference dulu dengan tanda *
	// kalau tidak maka yang di set hanyalah
	// memory address pointer *film
	*film = newData

	// newData akan terhapus dari memory
	// karena tidak ada variabel yang menampungnya
	// maka ketika di set ke pointer film
	// maka pointer film akan menunjuk ke memory
	// yang sebelumnya di pakai oleh newData
	// yaitu Film{} kosong
	// film = &newData
}

// Tanpa membuat variabel newData
func InputData2(film *Film) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Judul: ")
	film.Judul, _ = scanner.ReadString('\n')
	film.Judul = strings.TrimSpace(film.Judul)

	fmt.Print("Masukkan Rating: ")
	ratingStr, _ := scanner.ReadString('\n')
	ratingStr = strings.TrimSpace(ratingStr) // menghilangkan karakter '\n'
	newRating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	film.Rating = newRating

	fmt.Print("Masukkan Deskripsi : ")
	film.Deskripsi, _ = scanner.ReadString('\n')
	film.Deskripsi = strings.TrimSpace(film.Deskripsi)

	fmt.Print("Masukkan Studio : ")
	film.Studio, _ = scanner.ReadString('\n')
	film.Studio = strings.TrimSpace(film.Studio)

	fmt.Print("Masukkan durasi : ")
	durasiStr, _ := scanner.ReadString('\n')
	durasiStr = strings.TrimSpace(durasiStr)
	newDurasi, err := strconv.Atoi(durasiStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	film.Durasi = newDurasi

	fmt.Print("Masukkan Gendre : ")
	film.Gendre, _ = scanner.ReadString('\n')
	film.Gendre = strings.TrimSpace(film.Gendre)
}

func main() {
	// Metode pertama
	var film1 Film
	InputData1(&film1)
	fmt.Println(film1)
	fmt.Println("Film 1 berhasil di input")

	// Metode kedua
	var film2 Film
	InputData2(&film2)
	fmt.Println(film2)
	fmt.Println("Film 2 berhasil di input")
}
