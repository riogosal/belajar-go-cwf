package module

import (
	"app-palindrome/model"
	"fmt"
)

func DeleteData(dataP map[model.Palindrome]bool) {
	fmt.Println("Masukkan Kata : ")
	var search string
	fmt.Scanln(&search)

	if _, ok := dataP[model.Palindrome{Name: search}]; ok {
		delete(dataP, model.Palindrome{Name: search})
		fmt.Println("data anda sudah dihapus")
	} else {
		fmt.Println("data Anda Tidak ada !")
	}

}
