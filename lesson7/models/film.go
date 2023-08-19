package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Film struct {
	Judul     string
	Rating    float64
	Deskripsi string
	Studio    string
	Durasi    int
}

func (f *Film) ToCSV() string {
	return fmt.Sprintf("%s,%f,%s,%s,%d\n", f.Judul, f.Rating, f.Deskripsi, f.Studio, f.Durasi)
}

func FromCSV(line string) Film {
	hasilSplit := strings.Split(line, ",")

	rating, err := strconv.ParseFloat(hasilSplit[1], 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	durasi, err := strconv.Atoi(hasilSplit[4])
	if err != nil {
		fmt.Println("Error:", err)
	}

	return Film{
		Judul:     hasilSplit[0],
		Rating:    rating,
		Deskripsi: hasilSplit[2],
		Studio:    hasilSplit[3],
		Durasi:    durasi,
	}
}

// This function can't be used opt for the strings.Split() function
// func FromCSV(csv string) Film {
// 	var f Film
// 	fmt.Sscanf(csv, "%s,%f,%s,%s,%d\n", &f.Judul, &f.Rating, &f.Deskripsi, &f.Studio, &f.Durasi)
// 	return f
// }
