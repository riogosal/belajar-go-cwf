package fungsi

import (
	"belajar-go-cwf/PR/eser/crudmap/model"
	"bufio"
	"fmt"
	"os"
)

func Load(nama_file string, koleksi_film map[string]model.Film) {
	fileHandler, err := os.Open(nama_file)
	if err != nil {
		fmt.Println("File tidak ditemukan, membuat file baru")

		if _, err := os.Create(nama_file); err != nil {
			panic(err)
		}
		return
	}
	defer fileHandler.Close()

	file_scanner := bufio.NewScanner(fileHandler)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		film := model.StringToFilm(line)
		koleksi_film[film.Judul] = film
	}
}
