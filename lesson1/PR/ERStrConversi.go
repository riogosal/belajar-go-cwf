package main

import (
	"fmt"
	"strconv"
)

func main() {
	int := "1a23"
	integers, err := strconv.Atoi(int)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(integers)
}
