package module

import (
	"belajar-go-cwf/PR/irfan/updateDataCsv/model"
	"os"
)

func Save(nama_file string, dataMovie []model.Movie) {
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
	for _, v := range dataMovie {
		file_handler.WriteString(v.ToCsv())
	}
}
