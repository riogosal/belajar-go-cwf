package main

import "fmt"

type Print interface {
	printInfo()
}

type Food struct {
	Name  string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (f Food) printInfo() {
	fmt.Printf("Food  : %s Price : %.2f\n", f.Name, f.Price)
}

func (d Drink) printInfo() {
	fmt.Printf("Drink : %s  Price : %.2f\n", d.Name, d.Price)
}

func Makanan(name string, price float32) Print {
	return Food{Name: name, Price: price}
}

func Minuman(name string, price float32) Print {
	return Drink{Name: name, Price: price}
}

func main() {
	food1 := Makanan("Nasi Goreng", 12.000)
	food2 := Makanan("Mie Ayam", 10.000)
	drink1 := Minuman("Kopi", 14.000)
	drink2 := Minuman("Teh", 8.000)

	info := []Print{food1, food2, drink1, drink2}

	for _, item := range info {
		item.printInfo()
	}
}
