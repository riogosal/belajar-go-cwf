package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Film struct {
	Judul     string
	Deskripsi string
	Durasi    int
	// Label     []Label
}
type Label string

func (f Film) Tampilkan() string {
	return fmt.Sprintf(
		"%s\nDeskripsi: %s\nDurasi: %d\n",
		f.Judul,
		f.Deskripsi,
		f.Durasi,
	)
}

func StringToFilm(csv_line string) Film {
	raw := strings.Split(csv_line, ",")

	durasi, err := strconv.Atoi(raw[2])
	if err != nil {
		panic(err)
	}
	return Film{
		Judul:     raw[0],
		Deskripsi: raw[1],
		Durasi:    durasi,
	}
}

func (f Film) ToCSV() string {
	return fmt.Sprintf(
		"%s,%s,%d\n",
		f.Judul,
		f.Deskripsi,
		f.Durasi,
	)
}
