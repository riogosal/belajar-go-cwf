package module

import (
	"belajar-go-cwf/PR/irfan/updateDataCsv/model"
	"bufio"
	"fmt"
	"os"
)

func Load(DATA_FILE string, dataMovies *[]model.Movie) {
	f, err := os.Open(DATA_FILE)

	if err != nil {
		fmt.Println("file CSV tidak ada, anda telah dibuatkan file CSV")
		if _, err := os.Create(DATA_FILE); err != nil {
			panic(err)
		}

	}

	file_scanner := bufio.NewScanner(f)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		film := model.StringToFilm(line)
		*dataMovies = append(*dataMovies, film)
	}

}
