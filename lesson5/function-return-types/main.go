package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Judul             string
	Rating            float64
	Deskripsi, Studio string
	Durasi            int
	Gendre            string
}

func CreateNewMovieWithJudulAndRating(
	judul string,
	rating string,
) (Movie, error) {
	ratingFloat, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		return Movie{}, err
	}
	return Movie{
		Judul:     judul,
		Rating:    ratingFloat,
		Deskripsi: "No Deskripsi",
		Studio:    "Studio 45",
		Durasi:    120,
		Gendre:    "Adventure",
	}, nil
}

func main() {
	harry_potter, err := CreateNewMovieWithJudulAndRating("Harry potter", "8.9") // coba ubah rating menjadi "8.9a" untuk melihat error dalam aksi
	if err != nil {
		panic(err) // ketika terjadi error, program akan berhenti dan menampilkan error
	}
	fmt.Println(harry_potter)
}
