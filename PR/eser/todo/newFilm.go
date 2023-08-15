package main

import (
	"bufio"
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

func InputDataGenre(film *Film) {
	for {
		fmt.Print("Masukkan Genre atau Tekan (e) untuk exit: ")
		var genreFilm Genre
		// fmt.Println(film.Genres)
		fmt.Scanln(&genreFilm)
		if genreFilm == "e" {
			break
		} else {
			film.Genres = append(film.Genres, genreFilm)
		}
	}
}

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

	InputDataGenre(film)
}

func UpdateNewFilm(film *Film) {
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

	InputDataGenre(film)
}

func DataFilm(film *Film) {
	fmt.Println("Daftar Film favorit")
	fmt.Println("Judul     :", film.Judul)
	fmt.Println("Deskripsi :", film.Deskripsi)
	fmt.Println("Studio    :", film.Studio)
	fmt.Println("Durasi    :", film.Durasi)
	fmt.Println("Genre     :")

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
}

func UpdateGenre(film *Film) {
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

	if inputGenre < 1 || inputGenre > len(film.Genres) { // Kita buat validasi, kalau user input 0 itu tidak bisa, karena index sekarang dihitung dari 1
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
}

func main() {
	var myFilm Film
	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. Buat Film Favorit ")
		fmt.Println("2. Tampilkan Film ")
		fmt.Println("3. Update Film ")
		fmt.Println("4. Hapus ")
		fmt.Println("5. Update Genres ")
		fmt.Println("Pilih Menu (1,2,3,4), Ketik (e) untuk keluar")
		fmt.Print("Input Menu : ")

		var inputMenu string
		fmt.Scanln(&inputMenu)

		if inputMenu == "e" {
			fmt.Println("Menutup Aplikasi!")
			break
		} else if inputMenu == "1" {
			InputDataFilm(&myFilm)
		} else if inputMenu == "2" {
			DataFilm(&myFilm)
		} else if inputMenu == "3" {
			UpdateNewFilm(&myFilm)
		} else if inputMenu == "4" {
			myFilm = Film{}
		} else if inputMenu == "5" {
			UpdateGenre(&myFilm)
		}
	}
}
