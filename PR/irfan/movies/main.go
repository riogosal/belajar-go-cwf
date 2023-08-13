package main

import (
	"app-data-movies/entity"
	"fmt"
)

func main() {
	fmt.Println("Selamat datang di program film favorit")

	var createData []int16
	createData = append(createData, 1)

	dataMovies := entity.Movies{
		UUID: createData[0],
	}

	dataMovies.CetakMovies()
}
