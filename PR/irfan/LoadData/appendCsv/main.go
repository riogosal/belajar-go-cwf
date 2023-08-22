package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Movie struct {
	Id, Tahun                int16
	Judul, Deskripsi, Negara string
}

func (m Movie) ToCSV() string {
	return fmt.Sprintf("%d,%d,%s,%s,%s\n", m.Id, m.Tahun, m.Judul, m.Deskripsi, m.Negara)
}

func main() {
	data := []Movie{
		{Id: 1, Tahun: 2002, Judul: "AADC", Deskripsi: "Film remaja", Negara: "Indonesia"},
		{Id: 2, Tahun: 2010, Judul: "Iron man", Deskripsi: "Film Hero", Negara: "USA"},
	}

	path := "file.csv"

	f, err := os.OpenFile("../"+path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic("error" + err.Error())
	}

	defer f.Close()

	w := csv.NewWriter(f)
	for _, p := range data {
		f.WriteString(p.ToCSV())
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)

	}

	fmt.Println("insert data into csv successfully")
}
