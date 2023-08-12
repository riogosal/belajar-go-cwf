package entity

import "fmt"

type Movie struct {
	ID           int64
	Judul        string
	Deskripsi    string
	Studio       string
	Durasi       int
	Status       bool
	JadwalTayang []Jadwal
	GendreFilm   []Genre
}

func (movie Movie) Print() {
	fmt.Println("Id          : ", movie.ID)
	fmt.Println("Judul Film  : ", movie.Judul)
	fmt.Println("Deskripsi   : ", movie.Deskripsi)
	fmt.Println("Studio      : ", movie.Studio)
	fmt.Println("Durasi Film : ", movie.Durasi)
	fmt.Println("Status      : ", movie.Status)
	// fmt.Println("Jadwal      : ", movie.JadwalTayang)
	for _, movie := range movie.JadwalTayang {
		movie.PrintJadwal()
	}

	for _, movie := range movie.GendreFilm {
		movie.PrintGendre()
	}
	// fmt.Println("Gendre      : ", movie.GendreFilm)
}

// Ini sama dengan fungsi bawah dan tidak berguna
func (movie Movie) UbahJudul(judul string) {
	movie.Judul = judul
	fmt.Println("Dalam fungsi ubah judul", movie)
}
func UbahJudul(movie Movie, judul string) {
	movie.Judul = judul
}

// Cara yang benar untuk membuat fungsi mengubah movie
func (movie *Movie) UbahJudulPakaiPointer(judul string) {
	movie.Judul = judul
}

// Function sama dengan di atas
func UbahJudulPakaiPointer(movie *Movie, judul string) {
	movie.Judul = judul
}
