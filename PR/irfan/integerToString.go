package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int = 123
	nilai2 := strconv.Itoa(a)
	fmt.Printf("Type %T\n ", nilai2)
	fmt.Println("Nilainya ", nilai2)
}
