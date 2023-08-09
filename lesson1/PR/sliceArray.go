package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numbers []int // Variabel numbers digunakan untuk menampung nilai-nilai dalam bentuk slice

	for {
		fmt.Print("Masukkan nomor atau ketik 'exit' untuk selesai :")
		var input string
		fmt.Scan(&input)

		if input == "exit" {
			if len(numbers) == 0 {
				fmt.Println("Menutup Program, Terima Kasih")
				return
			}
			datainput(numbers)
			return
		}

		nomor, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Harap Masukkan Angka!")
			continue
		}
		numbers = append(numbers, nomor) // digunakan untuk menambahkan nilai 'nomor' ke dalam slice 'numbers'
		jumlah := len(numbers)
		fmt.Println(numbers)
		fmt.Println("Jumlah Nomor yang telah di input : ", jumlah)
	}
}

func datainput(numbers []int) {
	jumlah := len(numbers)
	fmt.Println("Jumlah Nomor yang telah di input : ", jumlah)
	fmt.Println("Nomor yang di input :", numbers)
	fmt.Println("Total : ", totaljumlah(numbers))
	fmt.Println("Nilai Rata-rata:", average(numbers))
	fmt.Println("Nilai Median:", median(numbers))
}

func totaljumlah(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func average(numbers []int) float32 {
	total := totaljumlah(numbers)
	return float32(total) / float32(len(numbers))
}

func median(numbers []int) float32 {
	length := len(numbers)

	if length%2 != 0 {
		sort.Ints(numbers) // Mengurutkan slice numbers dari angka terkecil ke terbesar

		dataindex := (length - 1) / 2
		nilai := float32(numbers[dataindex])

		return nilai

	} else {
		sort.Ints(numbers) // Mengurutkan slice numbers dari angka terkecil ke terbesar
		// fmt.Println(numbers)
		nilai_length1 := numbers[(length-1)/2]
		nilai_length2 := numbers[(length / 2)]
		hasil := (float32(nilai_length1) + float32(nilai_length2)) / 2

		return hasil
	}
	// 20, 28, 35, 37, 41, 43, 55, 58, 60, 63, 70, 74, 76, 84
	return float32(length)
}
