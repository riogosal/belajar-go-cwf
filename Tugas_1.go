// Convert fahrenheit to celcius By go
package main

import "fmt"

func convertTempToCelcius(fahrenheit float32) {

	var result = (fahrenheit - 32) * 5 / 9
	fmt.Println("nilai dari ", fahrenheit, "fahrenheit = ", result, "Celcius")
}

func main() {
	type temp float32
	var fahrenheit int = 250
	println(fahrenheit)
	convertTempToCelcius(float32(fahrenheit))
}
