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

func InputDataFilm() {
	var film Film
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
	fileJson := "data_json.json"

	/* ref :
	https://go.dev/src/os/example_test.go
	 https://webdamn.com/how-to-read-csv-file-using-golang/
	*/

	f, err := os.OpenFile(fileCsv, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

	fmt.Println("Berhasil menyimpan data CSV dengan nama File :", fileCsv)

	// ---- Simpan Data Ke json ----//

	// Ubah data ke dalam bentuk Json
	jsonData, err := json.Marshal(film)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	err = os.WriteFile(fileJson, jsonData, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Berhasil menympan data Json dengan nama File :", fileJson)

}

func main() {
	fileJson := "data_json.json"

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
			InputDataFilm()
		} else if inputMenu == "2" {

			jsonData, err := os.ReadFile(fileJson)

			if err != nil {
				fmt.Println("Belum ada file atau data")
				continue
			}
			fmt.Println(string(jsonData))

			var myDataJson *Film

			err = json.Unmarshal(jsonData, &myDataJson)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("\nData Json Yang Telah di parsing :")
			fmt.Println("Judul     :", myDataJson.Judul)
			fmt.Println("Rating    :", myDataJson.Rating)
			fmt.Println("Deskripsi :", myDataJson.Deskripsi)
			fmt.Println("Studio    :", myDataJson.Studio)
			fmt.Println("Durasi    :", myDataJson.Durasi)
		}
	}
}
