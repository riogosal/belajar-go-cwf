package fungsi

import (
	"belajar-go-cwf/PR/eser/crud/model"
	"bufio"
	"fmt"
	"os"
)

func Load(nama_file string, koleksi_film *[]model.Film) {
	fileHandler, err := os.Open(nama_file)
	if err != nil {
		fmt.Println("File Tidak di Temukan, Membuat File Baru")

		if err != nil {
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
