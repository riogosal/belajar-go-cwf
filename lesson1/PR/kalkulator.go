package main

import (
	"fmt"
	"strconv"
)

func main() {
	var angka1 string
	var angka2 string
<<<<<<< HEAD

=======
>>>>>>> a581ee45a48831c616c410ce9eef3560c12ef911
	var inputangka1 int
	var inputangka2 int

	var err error

	for {
<<<<<<< HEAD
		fmt.Println("\nPilihan Operator:")
=======
		fmt.Println("Pilihan Operasi:")
>>>>>>> a581ee45a48831c616c410ce9eef3560c12ef911
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
<<<<<<< HEAD
		fmt.Println("5. Exit")
		fmt.Print("Silahkan Pilih : ")

		var pilihan string

=======
		fmt.Println("5. Keluar")
		fmt.Print("Pilih operasi (1/2/3/4/5): ")
		var pilihan string
>>>>>>> a581ee45a48831c616c410ce9eef3560c12ef911
		fmt.Scanln(&pilihan)

		if pilihan == "5" {
			fmt.Println("Menutup kalkulator")
			break
		} else if pilihan != "1" && pilihan != "2" && pilihan != "3" && pilihan != "4" {
			fmt.Println("Menu yang di pilih tidak ada")
			continue
		}

<<<<<<< HEAD
		var operator string

		switch pilihan {
		case "1":
			operator = "Jumlah"
		case "2":
			operator = "Kurang"
		case "3":
			operator = "Perkalian"
		case "4":
			operator = "Pembagian"
		}

		fmt.Printf("Anda memilih: %s\n", operator)

=======
		fmt.Println("Menu", pilihan)
>>>>>>> a581ee45a48831c616c410ce9eef3560c12ef911
		fmt.Print("Angka Pertama: ")
		fmt.Scanln(&angka1)

		inputangka1, err = strconv.Atoi(angka1)

		if err != nil {
			fmt.Println("Mohon masukkan angka!")
			continue
		}

		fmt.Print("Angka Kedua  : ")
		fmt.Scanln(&angka2)

		inputangka2, err = strconv.Atoi(angka2)

		if err != nil {
			fmt.Println("Mohon masukkan angka!")
			continue
		}

		var hasil int
		if pilihan == "1" {
			hasil = inputangka1 + inputangka2
		} else if pilihan == "2" {
			hasil = inputangka1 - inputangka2
		} else if pilihan == "3" {
			hasil = inputangka1 * inputangka2
		} else if pilihan == "4" {
			if inputangka2 != 0 {
				hasil = inputangka1 / inputangka2
			} else {
				fmt.Println("Tidak bisa membagi dengan nol.")
				continue
			}
		}

		fmt.Printf("Hasil: %d\n", hasil)
	}
}
