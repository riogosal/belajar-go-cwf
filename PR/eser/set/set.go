package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := []int{5, 10, 15, 20, 25, 30, 40, 100}

	var inputData []int
	for {
		fmt.Print("Input Number, (e) exit : ")
		var userInput string
		fmt.Scan(&userInput)

		if userInput == "e" {
			if len(inputData) == 0 {
				fmt.Println("Bye!")
			} else {
				updateData(&data, inputData)
				fmt.Println("Data yang telah di update :", data)
			}
			return
		}

		angka, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Harap Masukkan Angka!")
			continue
		}
		inputData = append(inputData, angka)
	}
}

func updateData(data *[]int, inputData []int) {
	for _, myData := range inputData {
		found := false
		for _, number := range *data {
			if myData == number {
				found = true
				break
			}
		}

		if !found {
			*data = append(*data, myData)
		}

		if found {
			fmt.Printf("Angka %d ditemukan\n", myData)
		} else {
			fmt.Printf("Angka %d tidak ditemukan, menambah angka.\n", myData)
		}
	}
}
