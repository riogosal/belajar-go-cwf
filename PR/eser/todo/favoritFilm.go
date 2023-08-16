package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var DataInput []Film
	statuDataFilm := false

	for {
		fmt.Println("Menu: ")
		fmt.Println("1. Buat Film Favorit ")
		fmt.Println("2. Tampilkan Film ")
		fmt.Println("3.  Update Film ")
		fmt.Println("Pilih Menu (1,2), Ketik (e) untuk keluar")
		fmt.Print("Input Menu : ")

		var inputMenu string
		fmt.Scanln(&inputMenu)

		if inputMenu == "e" {
			fmt.Println("Menutup Aplikasi")
			break
		} else if inputMenu == "1" {
			if statuDataFilm == false {
				fmt.Println("Membuat Film Favorit")
				InputData(&DataInput)
				statuDataFilm = true
			}
			fmt.Println("Film favorit telah di input pilih menu (2) untuk melihat film")
		} else if inputMenu == "2" {
			if statuDataFilm {
				fmt.Println("Daftar Film Favorit Kamu :")
				for _, film := range DataInput {
					film.printFilm()
				}
			}

		} else if inputMenu == "3" {
			fmt.Println("Update Film Favorit")

			UpdateData(&DataInput)

			statuDataFilm = true
		} else if inputMenu != "1" && inputMenu != "2" && inputMenu != "3" {
			fmt.Println("Menu yang di pilih tidak ada!")
			continue
		}
	}
}

type Film struct {
	Judul             string
	Rating            float64
	Deskripsi, Studio string
	Durasi            int
	Gendre            string
}

func (film *Film) printFilm() {
	fmt.Println("Judul Film     : ", film.Judul)
	fmt.Println("Rating Film    : ", film.Rating)
	fmt.Println("Deskripsi Film : ", film.Deskripsi)
	fmt.Println("Studio Film    : ", film.Studio)
	fmt.Println("Durasi Film    : ", film.Durasi)
	fmt.Println("Gendre Film    : ", film.Gendre)
}
func UpdateData(DataInput *[]Film) {
	fmt.Println(&DataInput)
	fmt.Println(*DataInput)

<<<<<<< HEAD
func InputData(DataInput *Film) {
=======
	var judul string
	// scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Update Judul: ")
	// judul, _ = scanner.ReadString('\n')
	// judul = strings.TrimSpace(judul)

	for _, value := range *DataInput {
		fmt.Printf("%p", &value)
		value.Judul = judul
		fmt.Println(value.Judul)
	}

	// *DataInput * []Film.Judul = judul

}
func InputData(DataInput *[]Film) {
>>>>>>> 5615a72f8b392a70c785d112c84dcec7f3ec42b8
	var judul string
	var rating float64
	var deskripsi string
	var studio string
	var durasi int
	var gendre string

	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Judul: ")
	judul, _ = scanner.ReadString('\n')
	judul = strings.TrimSpace(judul)

	fmt.Print("Masukkan Rating: ")
	ratingStr, _ := scanner.ReadString('\n')
	ratingStr = strings.TrimSpace(ratingStr) // menghilangkan karakter '\n'
	newRating, err := strconv.ParseFloat(ratingStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rating = newRating

	fmt.Print("Masukkan Deskripsi : ")
	deskripsi, _ = scanner.ReadString('\n')
	deskripsi = strings.TrimSpace(deskripsi)

	fmt.Print("Masukkan Studio : ")
	studio, _ = scanner.ReadString('\n')
	studio = strings.TrimSpace(studio)

	fmt.Print("Masukkan durasi : ")
	durasiStr, _ := scanner.ReadString('\n')
	durasiStr = strings.TrimSpace(durasiStr)
	newDurasi, err := strconv.Atoi(durasiStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	durasi = newDurasi

	fmt.Print("Masukkan Gendre : ")
	gendre, _ = scanner.ReadString('\n')
	gendre = strings.TrimSpace(gendre)

	newData := Film{
		Judul:     judul,
		Rating:    rating,
		Deskripsi: deskripsi,
		Studio:    studio,
		Durasi:    durasi,
		Gendre:    gendre,
	}

	*DataInput = append(*DataInput, newData)
}
