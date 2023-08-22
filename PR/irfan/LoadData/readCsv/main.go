package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	Id, Tahun                int
	Judul, Deskripsi, Negara string
}

func (m Movie) ToCSV() string {
	return fmt.Sprintf("%d,%d,%s,%s,%s\n", m.Id, m.Tahun, m.Judul, m.Deskripsi, m.Negara)
}

func FromCSV(line string) Movie {
	hasilSplit := strings.Split(line, ",")

	id, err := strconv.Atoi(hasilSplit[0])
	if err != nil {
		fmt.Println("Error:", err)
	}

	tahun, err := strconv.Atoi(hasilSplit[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	return Movie{
		Id:        id,
		Judul:     hasilSplit[2],
		Deskripsi: hasilSplit[3],
		Negara:    hasilSplit[4],
		Tahun:     tahun,
	}
}

func (m *Movie) CetakMovies() {
	fmt.Println("ID FIlM : ", m.Id)
	fmt.Println(" JUDUL FILM : ", m.Judul)
	fmt.Println(" DESKRIPSI FILM : ", m.Deskripsi)
	fmt.Println(" ASAL NEGARA : ", m.Negara)
	fmt.Println(" TAHUN RILIS : ", m.Tahun)

}

func UpdateDataFilm(dataFilm *[]Movie) {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println(dataFilm)
	fmt.Println("Masukkan ID Film ")
	var update int
	fmt.Scanln(&update)
	for _, v := range *dataFilm {
		fmt.Print("Masukkan Judul: ")
		v.Judul, _ = scanner.ReadString('\n')
		v.Judul = strings.TrimSpace(v.Judul)

		fmt.Println(v)
		break
		// 	if v.Id == update {
		// fmt.Println(dataFilm[i])
		// }
	}
}

func main() {

	// var data Movie
	data := []Movie{}

	path := "file.csv"

	f, err := os.Open("../" + path)

	if err != nil {
		panic("error " + err.Error())
	}

	defer f.Close()
	// read csv values using csv.Reader
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		line := reader.Text()

		// do something with read line
		data = append(data, FromCSV(line))
	}

	fmt.Println(data[0])
	fmt.Println("Data from Csv to Program successfully read")
	for {
		fmt.Println("Options: 1. Melihat data, 2 Update Data, 3 Save & Exit...")
		var option string
		fmt.Scanln(&option)

		switch option {
		case "1":
			for _, v := range data {
				fmt.Println("<================================> Film ", v.Id)
				v.CetakMovies()
			}

		case "2":
			UpdateDataFilm(&data)

		}
	}

}
