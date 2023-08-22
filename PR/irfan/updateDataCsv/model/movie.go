package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Movie struct {
	Id, Tahun        int
	Judul, Deskripsi string
}

func (f Movie) CetakFilm() string {
	return fmt.Sprintf(
		"%d\nId :%s\nJudul: %s\nDeskripsi: %d\nTahun :",
		f.Id,
		f.Judul,
		f.Deskripsi,
		f.Tahun,
	)
}

func (m Movie) TampilFilm() {
	fmt.Println("================================")
	fmt.Println("ID : ", m.Id)
	fmt.Println("JUDUL FILM : ", m.Judul)
	fmt.Println("DESKRIPSI FILM : ", m.Deskripsi)
	fmt.Println("TAHUN FILM : ", m.Tahun)
}

func StringToFilm(line string) Movie {
	dataFilm := strings.Split(line, ",")

	id, err := strconv.Atoi(dataFilm[0])
	if err != nil {
		panic("error" + err.Error())

	}

	tahun, err := strconv.Atoi(dataFilm[3])
	if err != nil {
		panic("error" + err.Error())
	}

	return Movie{
		Id:        id,
		Judul:     dataFilm[1],
		Deskripsi: dataFilm[2],
		Tahun:     tahun,
	}
}

func (m Movie) ToCsv() string {
	return fmt.Sprintf("%d,%s,%s,%d\n", m.Id, m.Judul, m.Deskripsi, m.Tahun)

}
