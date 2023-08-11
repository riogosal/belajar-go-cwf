package main

import "fmt"

type Animal struct {
	JenisSpesis string
	Umur        int
}

func main() {
	singa := Animal{
		JenisSpesis: "Singa",
		Umur:        2,
	}
	fmt.Println(singa)

	tuna := Animal{
		JenisSpesis: "Thunnus Albacares",
		Umur:        1,
	}
	fmt.Println(tuna)

	// Accessing property of struct
	umurSinga := singa.Umur
	fmt.Println("Umur singa ini adalah", umurSinga)

	// Put struct into collection
	kebunBinatang := []Animal{}
	fmt.Println(kebunBinatang)
	fmt.Println(len(kebunBinatang))

	// Appending Animal struct to collection of Animals / array of Animals
	kebunBinatang = append(kebunBinatang, singa, tuna)
	fmt.Println(kebunBinatang)
}
