package main

import "fmt"

const NOMOR_FAVORIT_SAYA = 180

func main() {
	// Array integer
	x := []int{
		1, 2, 3, 4,
	}
	fmt.Println(x)

	// Array string
	arrayString := []string{
		"", "a",
	}
	fmt.Println(arrayString)

	// Menggunakan index di array
	test2 := arrayString[1]
	fmt.Println(test2)

	// Membuat mutasi pada arrayString
	arrayString[0] = "b"
	fmt.Println(arrayString)

	// Array vs Slice
	arrayString = append(arrayString, "d")
	fmt.Println(arrayString)

	// Slice
	arrayKosong := []int{}
	fmt.Println(arrayKosong)
	arrayKosong = append(arrayKosong, 100)
	arrayKosong = append(arrayKosong, 101)
	arrayKosong = append(arrayKosong, 102)
	fmt.Println(arrayKosong)

	// Array
	realArray := [4]int{}
	realArray[0] = 100
	realArray[1] = 101
	realArray[2] = 102
	fmt.Println(realArray)

	// Slice yang berbentuk array
	arrayLikeSlice := make([]int, 4)
	fmt.Println(arrayLikeSlice)
}
