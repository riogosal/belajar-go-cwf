package main

import (
	"fmt"
	"lesson4/entity"
)

func main() {
	movie1 := entity.Movie{
		ID:        1,
		Judul:     "Iron Man 3",
		Deskripsi: "Film action yang di perankan oleh Downey Jr",
		Studio:    "Marvel",
		Durasi:    123,
		Status:    false,
		JadwalTayang: []entity.Jadwal{
			{
				Jam:    "17.00",
				Tempat: "Mtos",
			},
			{
				Jam:    "18.00",
				Tempat: "Mtos",
			},
			{
				Jam:    "17.00",
				Tempat: "Nipah Mall",
			},
		},
		GendreFilm: []entity.Genre{
			"Action", "Romance",
		},
	}

	// Fungsi tidak berguna
	entity.UbahJudul(movie1, "Iron Man 2")
	movie1.UbahJudul("Iron Man 2")
	// Yang tidak merubah movie1
	fmt.Println(movie1)

	// Sama juga dengan di atas fungsi ini tidak merubah movie1 karena hanya merubah copyan
	// Tetapi dengan fungsi yang bukan bawahan Movie
	entity.UbahJudul(movie1, "Iron Man 2")
	fmt.Println(movie1)

	// Pengunaan fungsi struct yang benar
	movie1.UbahJudulPakaiPointer("Iron Man 1")
	fmt.Println(movie1)

	// Masukkan memory address movie1 kedalam fungsi berarti kita membuat link
	// bukan copyan seperti fungsi entity.UbahJudul(movie, "Judul")
	entity.UbahJudulPakaiPointer(&movie1, "Iron Man 100")
	fmt.Println(movie1)

	movie2 := entity.Movie{
		ID:        2,
		Judul:     "Superman",
		Deskripsi: "Film Action yang di perankan oleh Henry Cavill",
		Studio:    "DC",
		Durasi:    160,
		Status:    true,
		JadwalTayang: []entity.Jadwal{
			{
				Jam:    "17.00",
				Tempat: "Mtos",
			},
			{
				Jam:    "18.00",
				Tempat: "Mtos",
			},
			{
				Jam:    "17.00",
				Tempat: "Nipah Mall",
			},
		},
		GendreFilm: []entity.Genre{
			"Action",
		},
	}

	// Contoh menambah genre
	newGenre := entity.Genre("Fiction")

	// ------------ TANPA POINTER ----------------
	// Membuat array yang copy movie1 dan movie2
	daftarFilm := [...]entity.Movie{movie1, movie2}

	fmt.Printf("Memory address film1: %p\n", &movie1)
	fmt.Printf("Value index 0: %v\n", daftarFilm[0])
	fmt.Printf("Memory address index 0: %p\n", &daftarFilm[0])
	fmt.Println()
	fmt.Printf("Memory address film2: %p\n", &movie2)
	fmt.Printf("Value index 1: %v\n", daftarFilm[1])
	fmt.Printf("Memory address index 1: %p\n", &daftarFilm[1])

	for _, film := range daftarFilm {
		film.GendreFilm = append(film.GendreFilm, newGenre)
	}

	// movie1 dan movie2 tidak berubah karena daftarFilm[0] adalah copyan dari movie1, begitu juga dengan movie2
	fmt.Println(movie1.GendreFilm)
	fmt.Println(movie2.GendreFilm)

	// ------------ PAKAI POINTER ----------------
	// Membuat array dengan index yang link movie1 dan movie2
	daftarFilmPakaiPointer := [...]*entity.Movie{&movie1, &movie2}

	fmt.Printf("Memory address film1: %p\n", &movie1)
	fmt.Printf("Value index 0: %p\n", daftarFilmPakaiPointer[0])
	fmt.Printf("Memory address index 0: %p\n", &daftarFilmPakaiPointer[0])
	fmt.Println()
	fmt.Printf("Memory address film2: %p\n", &movie2)
	fmt.Printf("Value index 1: %p\n", daftarFilmPakaiPointer[1])
	fmt.Printf("Memory address index 1: %p\n", &daftarFilmPakaiPointer[1])

	for _, film := range daftarFilmPakaiPointer {
		film.GendreFilm = append(film.GendreFilm, newGenre)
	}

	fmt.Println(movie1.GendreFilm)
	fmt.Println(movie2.GendreFilm)
}
