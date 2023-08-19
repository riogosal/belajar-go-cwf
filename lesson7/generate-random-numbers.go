package main

import (
	"fmt"
	"math/rand"
	"os"
)

func GenerateRandomNumbersInAFile() {
	// Seed the random number generator
	// rand.Seed(time.Now().UnixNano())

	// Open the file for writing
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define the range of random numbers
	min := 100000000
	max := 100000000000

	// Number of random numbers to generate
	numNumbers := 1000000

	// Generate and write random numbers to the file
	for i := 0; i < numNumbers; i++ {
		randomNumber := rand.Intn(max-min+1) + min
		_, err := file.WriteString(fmt.Sprintf("%d\n", randomNumber))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Random numbers written to random_numbers.txt")
}
