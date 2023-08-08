package main

import "fmt"

func convertTempToCelcius(fahrenheit float32) {

	var result = (fahrenheit - 32) * 5 / 9
	fmt.Println("nilai dari ", fahrenheit, "fahrenheit = ", result, "Celcius")
}

func main() {
	type temp float32
	var fahrenheit temp = 250
	convertTempToCelcius(float32(fahrenheit))
}
