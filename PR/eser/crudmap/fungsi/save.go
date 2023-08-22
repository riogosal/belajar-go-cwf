package fungsi

import (
	"belajar-go-cwf/PR/eser/crudmap/model"
	"os"
)

func Save(nama_file string, koleksi_film map[string]model.Film) {
	file_hendler, err := os.OpenFile(nama_file, os.O_WRONLY, 0644)

	defer func() {
		err := file_hendler.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)

	}
	for _, film := range koleksi_film {
		file_hendler.WriteString(film.ToCSV())
	}
}
