package fungsi

import (
	"belajar-go-cwf/lesson8/model"
	"bufio"
	"fmt"
	"os"
)

func Load(nama_file string, koleksi_film *[]model.Film) {
	fileHandler, err := os.Open(nama_file)
	if err != nil {
		fmt.Println("No files are found, creating one")

		if _, err := os.Create(nama_file); err != nil {
			panic(err)
		}
	}

	file_scanner := bufio.NewScanner(fileHandler)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		film := model.StringToFilm(line)
		*koleksi_film = append(*koleksi_film, film)
	}
}
