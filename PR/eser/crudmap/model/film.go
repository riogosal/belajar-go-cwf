package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Film struct {
	ID        int
	Judul     string
	Deskripsi string
	Studio    string
	Rating    float64
	Durasi    int
}

func (f Film) Tampilkan() string {
	return fmt.Sprintf(
		"ID	  : %d\nJudul     : %s\nDeskripsi : %s\nStudio    : %s\nRating    : %f\nDurasi    : %d\n",
		f.ID,
		f.Judul,
		f.Deskripsi,
		f.Studio,
		f.Rating,
		f.Durasi,
	)
}

func StringToFilm(csv_line string) Film {
	dataCsv := strings.Split(csv_line, ",")
	// fmt.Println(dataCsv)5
	ID_int, err := strconv.Atoi(dataCsv[0])
	if err != nil {
		panic(err)
	}
	ratingFloat, err := strconv.ParseFloat(dataCsv[4], 64)
	if err != nil {
		panic(err)
	}

	durasiStr, err := strconv.Atoi(dataCsv[5])
	if err != nil {
		panic(err)
	}

	return Film{
		ID:        ID_int,
		Judul:     dataCsv[1],
		Deskripsi: dataCsv[2],
		Studio:    dataCsv[3],
		Rating:    ratingFloat,
		Durasi:    durasiStr,
	}
}

func (f Film) ToCSV() string {
	return fmt.Sprintf(
		"%d,%s,%s,%s,%f,%d\n",
		f.ID,
		f.Judul,
		f.Deskripsi,
		f.Studio,
		f.Rating,
		f.Durasi,
	)
}
