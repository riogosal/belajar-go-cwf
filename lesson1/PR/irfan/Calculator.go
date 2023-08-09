package main

import (
	"fmt"
	"time"
)

func ProgramEnd() {
	now := time.Now()
	fmt.Println("program Kalkulator By Muhammad Irfan Selesai, ", now.Day(), " - ", now.Month(), " - ", now.Year(), ". Pukul ", now.Hour(), " : ", now.Minute())
}

func Kalkulator(nilaiA int, nilaiB int, operator string) {
	defer ProgramEnd()
	fmt.Println("Nilai A :", nilaiA)
	fmt.Println("Nilai B :", nilaiB)

	switch operator {
	case "1":
		fmt.Println("hasilnya dikalikan :", nilaiA*nilaiB)
	case "2":
		fmt.Println("hasilnya dibagi :", float32(nilaiA)/float32(nilaiB))
	case "3":
		fmt.Println("hasilnya ditambahkan :", nilaiA+nilaiB)
	case "4":
		fmt.Println("hasilnya dikurang :", nilaiA-nilaiB)
	default:
		panic("error: anda salah memasukkan operator salah : " + operator)

	}

}

func main() {
	type value int16
	type cal string
	fmt.Println("Go runs Program Kalkulator On ")

	fmt.Println("Masukkan Nilai A :")
	var nilaiA value
	fmt.Scanln(&nilaiA)

	fmt.Println("Masukkan Nilai B :")
	var nilaiB value
	fmt.Scanln(&nilaiB)

	fmt.Println("Masukkan Oerator Mate-Matika : 1.kali, 2.bagi, 3.tambah, 4.kurang")
	var operator cal
	fmt.Scanln(&operator)

	Kalkulator(int(nilaiA), int(nilaiB), string(operator))

}
