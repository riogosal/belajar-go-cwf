package main

import (
	"fmt"
	"strconv"
)

// konversi nilai String to Integers
func StrToInt() {
	str := "123"
	// strconv.Atoi untuk mengkonversi string menjadi integer.
	// Fungsi ini mengembalikan dua nilai, hasil konversi dan error (jika ada error).
	result, err := strconv.Atoi(str)
	//Memeriksa apakah ada error saat melakukan konversi nilai str.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hasil Konversi String To Integer :", result)
}

// konversi nilai integer to string
func IntToStr() {
	int := 3043
	// Mengkonversi nilai integer menjadi string
	stringVersion := strconv.Itoa(int)
	fmt.Println("Hasil Konversi Integer To String :", stringVersion)
}

func main() {
	StrToInt()
	IntToStr()
}
