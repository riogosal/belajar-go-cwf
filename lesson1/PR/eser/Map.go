package main

import "fmt"

func main() {
	animal := map[string]int{
		"Lion":    2,
		"Chicken": 10,
		"Buffalo": 2,
		"Cat":     12,
	}
	animal["Cat"] = 89
	fmt.Println(animal)
	fmt.Printf("Jumlah Kucing Saya %d Dan Ayam %d Ekor", animal["Cat"], animal["Chicken"])
}
