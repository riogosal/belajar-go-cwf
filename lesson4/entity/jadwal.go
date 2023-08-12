package entity

import "fmt"

type Jadwal struct {
	Jam    string
	Tempat string
}

func (j Jadwal) PrintJadwal() {
	fmt.Println("Jam Tayang    : ", j.Jam)
	fmt.Println("Tempat Tayang : ", j.Tempat)
}
