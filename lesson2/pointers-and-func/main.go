package main

// func sayHello() {
// 	fmt.Println("Hello")
// }

// func sayHello2() {
// 	fmt.Println("Ello")
// }

func tambah1(nomor int) int {
	nomor = nomor + 1
	return nomor
}

// func tambah1_pointer(nomor *int) {
// 	*nomor = *nomor + 1
// }

func main() {

	// var pointer_x *int
	// fmt.Println(*pointer_x)

	// fmt.Println(pointer_x)
	// x := 4
	// x := 4
	// var pointer_untuk_x *int
	// pointer_untuk_x = &x

	nomorku := 0
	for x := int64(0); x < 1_000_000_000; x++ {
		nomorku = tambah1(nomorku)
	}

	// ekspektasi 16
}
