package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Genre string

type Film struct {
	Judul             string
	Rating            float64
	Deskripsi, Studio string
	Durasi            int
	Genres            []Genre
}

func removeGenre(film *Film, hapusGenre Genre) {
	for i, genre := range film.Genres {
		fmt.Println(i)
		if genre == hapusGenre {
			film.Genres = append(film.Genres[:i], film.Genres[i+1:]...)
			break // Keluar dari loop setelah elemen ditemukan dan dihapus
		}
	}
}

// func InputDataGenre(film *Film) {
// 	for {
// 		fmt.Print("Masukkan Genre atau Tekan (e) untuk exit: ")
// 		var genreFilm Genre
// 		// fmt.Println(film.Genres)
// 		fmt.Scanln(&genreFilm)
// 		if genreFilm == "e" {

// 			break
// 		} else {
// 			film.Genres = append(film.Genres, genreFilm)
// 		}
// 	}
// }

func InputDataFilm(film *Film) {
	// newData := Film{}

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

	// InputDataGenre(film)
}

func UpdateNewFilm(film *Film) {
	if film.Judul != "" {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Println("Silahkan pilih menu, 1. Update Judul, 2. Update Rating 3.Update Deskripsi 4. Update Studio, 5.Update Durasi, 6.Update Genre, (e) exit")
			var updateData string
			fmt.Scanln(&updateData)
			switch updateData {
			case "1":
				fmt.Print("Masukkan Judul Baru: ")
				film.Judul, _ = scanner.Text()
				film.Judul = strings.TrimSpace(film.Judul)
				fmt.Println("Judul berhasil diubah:", film.Judul)
				return
			case "2":
				fmt.Print("Masukkan Rating: ")
				ratingStr, _ := scanner.Text()
				ratingStr = strings.TrimSpace(ratingStr) // menghilangkan karakter
				newRating, err := strconv.ParseFloat(ratingStr, 64)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				film.Rating = newRating
				fmt.Println("Rating berhasil diubah:", film.Rating)
				return
			case "3":
				fmt.Print("Masukkan Deskripsi : ")
				film.Deskripsi, _ = scanner.Text()
				film.Deskripsi = strings.TrimSpace(film.Deskripsi)
				fmt.Println("Deskripsi berhasil diubah:", film.Deskripsi)
				return
			case "4":
				fmt.Print("Masukkan Studio : ")
				film.Studio, _ = scanner.Text()
				film.Studio = strings.TrimSpace(film.Studio)
				fmt.Println("Studio berhasil diubah:", film.Studio)
				return
			case "5":
				fmt.Print("Masukkan durasi : ")
				durasiStr, _ := scanner.Text()
				durasiStr = strings.TrimSpace(durasiStr)
				newDurasi, err := strconv.Atoi(durasiStr)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				film.Durasi = newDurasi
				fmt.Println("Judul berhasil Durasi:", film.Durasi)
				return
			case "6":
				UpdateGenre(film)
				return
			case "e":
				return
			}
		}
	} else {
		fmt.Println("Kamu belum memiliki data Film Favorit, Silahkan Input Film Terlebih dahulu")
	}
}

func DataFilm(film *Film) {
	if film.Judul != "" {
		fmt.Println("Daftar Film favorit")
		fmt.Println("Judul     :", film.Judul)
		fmt.Println("Deskripsi :", film.Deskripsi)
		fmt.Println("Studio    :", film.Studio)
		fmt.Println("Durasi    :", film.Durasi)
		fmt.Print("Genre     : ")
		length := len(film.Genres) - 1 // angaplah panjangnya 2
		for i, dataGenre := range film.Genres {
			fmt.Print(dataGenre)
			//pada saat ku loop, akan diperiksa apakah elemen tersebut adalah elemen terakhir, membandingkan index i dengan length (2)
			//Jika elemen tersebut bukan elemen terakhir, maka tambahkan koma.
			//Jika elemen tersebut adalah elemen terakhir, jangan tambahkan koma
			if i < length {
				fmt.Print(", ")
			}
		}
	} else {
		fmt.Println("Kamu belum memiliki daftar film Favorit, Silahkan Input Film Terlebih Dahulu")
	}
}

func UpdateGenre(film *Film) {
	if len(film.Genres) > 0 {
		var inputGenre int
		fmt.Println("Silakan pilih nomor yang ingin diubah:")
		for i, genre := range film.Genres {
			length := i + 1                       // kita tambah 1 karena index [film.Genres] itu mulai dari 0
			fmt.Printf("%d. %s\n", length, genre) //kita pakai %d karena length itu integer, kalo %s itu string
		}
		_, err := fmt.Scanln(&inputGenre) // kita parsing dalu data yang di input dari string ke integers
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if inputGenre < 1 || inputGenre > len(film.Genres) { // Kita buat validasi, kalau user input 0 itu tidak bisa, ///karena index sekarang dihitung dari 1
			fmt.Println("Nomor tidak valid") // kita hitung len (film.genders) lalu kita bandingkan dgn data yng di input, Intinya tidak boleh melebihi nilai Len
			return
		}
		var newGenre Genre // ini untuk variabel untuk menampung nilai dari &newGenre
		fmt.Print("Masukkan genre baru: ")
		_, err = fmt.Scanln(&newGenre)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		inputGenre = inputGenre - 1 // kita kurangi 1 karena di atas di tambahkan 1, biar dia membaca index mulai dari 0 dst
		film.Genres[inputGenre] = newGenre
		fmt.Println("Genre berhasil diubah:", newGenre)
	} else {
		fmt.Println("Anda Belum menginput data genre!")
	}
}

func main() {
	var myFilm Film
	namaFile := "film.json"

	// myFilmGenres := myFilm.Genres
	// fmt.Println(myFilmGenres)

	for {
		fmt.Println("\nMenu: ")
		fmt.Println("0. Lihat Data Json ")
		fmt.Println("1. Buat Film Favorit ")
		fmt.Println("2. Tampilkan Film ")
		fmt.Println("3. Update Film ")
		fmt.Println("4. Hapus ")
		fmt.Println("5. Update Genre ")
		fmt.Println("Pilih Menu (1,2,3,4,5), Ketik (e) untuk keluar")
		fmt.Print("Input Menu : ")

		var inputMenu string
		fmt.Scanln(&inputMenu)

		if inputMenu == "e" {
			fmt.Println("Menutup Aplikasi!")
			break
		} else if inputMenu == "0" {

			jsonData, err := os.ReadFile(namaFile)

			if err != nil {
				fmt.Println("Error:", err)
				return
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
			fmt.Println("Genres    :", myDataJson.Genres)

		} else if inputMenu == "1" {
			InputDataFilm(&myFilm)
			/*  Referensi:
			JSON Marshal Struct (Mengubah data ke Json):  https://golang.cafe/blog/golang-json-marshal-example.html
			Membaca data json : https://www.golangprograms.com/golang-writing-struct-to-json-file.html
			https://stackoverflow.com/questions/24770403/write-struct-to-json-file-using-struct-fields-not-json-keys

			All In 0ne : https://www.developer.com/languages/json-files-golang/
			*/

			// Ubah data ke dalam bentuk Json
			jsonData, err := json.Marshal(myFilm)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println(string(jsonData)) // kita tambahkan string untuk mengubah data byte ke string
			// Simpan data json dengan nama file "data.json"

			err = os.WriteFile(namaFile, jsonData, 0644)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Berhasil menympan data Json dengan nama File : film.json")

		} else if inputMenu == "2" {
			DataFilm(&myFilm)
		} else if inputMenu == "3" {
			UpdateNewFilm(&myFilm)
		} else if inputMenu == "4" {
			myFilm = Film{}
		}
		// else if inputMenu == "5" {
		// 	UpdateGenre(&myFilm)
		// }
	}
}
