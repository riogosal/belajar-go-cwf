package fungsi

import (
	"belajar-go-cwf/lesson8/model"
	"os"
)

func Save(nama_file string, koleksi_film []model.Film) {
	file_handler, err := os.OpenFile(nama_file, os.O_WRONLY, 0644)
	// Anonymous function
	defer func() {
		err := file_handler.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	for _, film := range koleksi_film {
		file_handler.WriteString(film.ToCSV())
	}
}
