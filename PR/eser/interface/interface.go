package main

import "fmt"

type Food struct {
	Name  string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

// type printer interface {
// 	printInfo()
// }

func (d Drink) printInfo() {
	fmt.Printf("Drink : %s  Price : %.2f\n", d.Name, d.Price)
}

func (b Food) printInfo() {
	fmt.Printf("Food  : %s Price : %.2f\n", b.Name, b.Price)
}

func main() {
	nasi_goreng := Food{
		Name:  "Nasi Goreng Jaksel",
		Price: 12.000,
	}

	kopi := Drink{
		Name:  "Kopi Gula Aren",
		Price: 14.000,
	}

	// nasi_goreng.printInfo()
	// kopi.printInfo()

	info := []printer{nasi_goreng, kopi}
	for _, item := range info {
		item.printInfo()
	}
}
