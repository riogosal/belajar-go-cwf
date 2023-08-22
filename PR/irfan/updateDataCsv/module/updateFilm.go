package module

import (
	"belajar-go-cwf/PR/irfan/updateDataCsv/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UpdateFilm(dataMovie []model.Movie) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Updating film....")
	fmt.Println("masukkan Id Film : ")
	var pilihanId int
	fmt.Scanln(&pilihanId)

	for i, v := range dataMovie {
		if v.Id == pilihanId {
			fmt.Println("pilih 1. Ubah Judul, 2. Ubah Deskripsi, 3. Ubah Tahun")
			var choiceUpdate int
			fmt.Scanln(&choiceUpdate)

			switch choiceUpdate {
			case 1:
				fmt.Print("Update Judul: ")
				scanner.Scan()
				judul := scanner.Text()
				dataMovie[i].Judul = judul
			case 2:
				fmt.Print("Update Deskripsi Film: ")
				scanner.Scan()
				desc := scanner.Text()
				dataMovie[i].Deskripsi = desc
			case 3:
				fmt.Print("Tahun: ")
				scanner.Scan()
				data_film_tahun := scanner.Text()
				tahun_int, err := strconv.Atoi(data_film_tahun)
				if err != nil {
					panic(err)
				}
				dataMovie[i].Tahun = tahun_int

			}

		}
	}

}
