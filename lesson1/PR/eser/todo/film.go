package main

import (
	"project-todo/entity"
)

func main() {
	film1 := entity.Movie{
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
		GendreFilm: []entity.Gendre{
			{
				Nama: "Action",
			},
			{
				Nama: "Romantic",
			},
		},
	}

	film2 := entity.Movie{
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
		GendreFilm: []entity.Gendre{
			{
				Nama: "Action",
			},
		},
	}

	daftarFilm := [...]entity.Movie{film1, film2}
	for _, film := range daftarFilm {
		film.PrintStuff()
	}
}

// type Movie struct {
// 	ID           int64
// 	Judul        string
// 	Deskripsi    string
// 	Studio       string
// 	Durasi       int
// 	Status       bool
// 	JadwalTayang []Jadwal
// 	GendreFilm   []Gendre
// }

// type Jadwal struct {
// 	Jam    string
// 	Tempat string
// }

// type Gendre struct {
// 	Nama string
// }

// func (j Jadwal) PrintJadwal() {
// 	fmt.Println("Jam Tayang    : ", j.Jam)
// 	fmt.Println("Tempat Tayang : ", j.Tempat)
// }

// func (g Gendre) PrintGendre() {
// 	fmt.Println("Gendre      : ", g.Nama)
// }

// func (movie Movie) PrintStuff() {
// 	fmt.Println("Id          : ", movie.ID)
// 	fmt.Println("Judul Film  : ", movie.Judul)
// 	fmt.Println("Deskripsi   : ", movie.Deskripsi)
// 	fmt.Println("Studio      : ", movie.Studio)
// 	fmt.Println("Durasi Film : ", movie.Durasi)
// 	fmt.Println("Status      : ", movie.Status)
// 	for _, movie := range movie.JadwalTayang {
// 		movie.PrintJadwal()
// 	}

// 	for _, movie := range movie.GendreFilm {
// 		movie.PrintGendre()
// 	}
// }
