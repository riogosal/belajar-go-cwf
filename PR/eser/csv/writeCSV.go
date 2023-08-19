package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Film struct {
	Judul     string
	Rating    float64
	Deskripsi string
	Studio    string
	Durasi    int
}

func (f *Film) ToCSV() string {
	return fmt.Sprintf("%s,%f,%s,%s,%d\n", f.Judul, f.Rating, f.Deskripsi, f.Studio, f.Durasi)
}

func FromCSV(csv string) Film {
	var f Film
	fmt.Sscanf(csv, "%s,%f,%s,%s,%d\n", &f.Judul, &f.Rating, &f.Deskripsi, &f.Studio, &f.Durasi)
	return f
}

func InputDataFilm(film *Film) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Judul: ")
	film.Judul, _ = scanner.ReadString('\n')
	film.Judul = strings.TrimSpace(film.Judul)

	fmt.Print("Masukkan Rating: ")
	ratingStr, _ := scanner.ReadString('\n')
	ratingStr = strings.TrimSpace(ratingStr)
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

	fileCsv := "data_film.csv"
	fileJson := "data_film.json"

	f, err := os.Create(fileCsv)

	if err != nil {
		fmt.Println("Gagal membuka File:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(film.ToCSV())
	if err != nil {
		fmt.Println("Gagal menulis data ke dalam file :", err)
		return
	}

	jsonData, err := json.Marshal(film)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData)) // kita tambahkan string untuk mengubah data byte ke string
	// Simpan data json dengan nama file "data.json"

	err = os.WriteFile(fileJson, jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Berhasil menympan data Json dengan nama File : ", fileJson)
}

func main() {

	var myFilm Film

	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. Buat Film Favorit ")
		fmt.Println("2. Lihat Data Json ")
		fmt.Println("Pilih Menu (1,2) Ketik (e) untuk keluar")
		fmt.Print("Input Menu : ")

		var inputMenu string
		fmt.Scanln(&inputMenu)

		if inputMenu == "e" {
			fmt.Println("Menutup Aplikasi!")
			break
		} else if inputMenu == "1" {
			InputDataFilm(&myFilm)

		} else if inputMenu == "4" {
			myFilm = Film{}
		}
	}

}
