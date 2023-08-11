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
	GendreFilm   []Gendre
}

type Jadwal struct {
	Jam    string
	Tempat string
}

type Gendre struct {
	Nama string
}

func (j Jadwal) PrintJadwal() {
	fmt.Println("Jam Tayang    : ", j.Jam)
	fmt.Println("Tempat Tayang : ", j.Tempat)
}

func (g Gendre) PrintGendre() {
	fmt.Println("Gendre      : ", g.Nama)
}

func (movie Movie) PrintStuff() {
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
