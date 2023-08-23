package module

import (
	"app-palindrome/model"
	"fmt"
)

func CariData(dataP map[model.Palindrome]bool) {
	fmt.Println("Masukkan Kata : ")
	var search string
	fmt.Scanln(&search)

	if val, ok := dataP[model.Palindrome{Name: search}]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("data Anda Tidak ada !")
	}

}
