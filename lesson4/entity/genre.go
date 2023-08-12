package entity

import "fmt"

// Karena sebelumnya hanya mempunyai satu tipe data
// saya ubah saja menjadi buatkan alias untuk string: Genre
//
//	type Genre struct {
//		Nama string
//	}

//   |
//   |
//   V

type Genre string

func (g Genre) PrintGendre() {
	fmt.Println("Gendre      : ", string(g))
}
