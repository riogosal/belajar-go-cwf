package main

import (
	"app-palindrome/model"
	"app-palindrome/module"
	"fmt"
)

func main() {
	dataPalindrome := make(map[model.Palindrome]bool)
close:
	for {
		fmt.Println("Pilih Menu : 1. Input Data, 2. Cari Data, 3. Hapus data, 4. Tampilkan data, 5. Exit ")
		var menu string
		fmt.Scanln(&menu)

		switch menu {
		case "1":
			module.CreateData(dataPalindrome)
		case "2":
			module.CariData(dataPalindrome)
		case "3":
			module.DeleteData(dataPalindrome)
		case "4":
			fmt.Println(dataPalindrome)
		default:
			print("terima kasih, anda hebat !")
			break close

		}

	}

}
