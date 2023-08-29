package main

import "fmt"

type printer interface {
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

func createFood(name string, price float32) printer {
	return Food{Name: name, Price: price}
}

func createDrink(name string, price float32) printer {
	return Drink{Name: name, Price: price}
}

func main() {
	food1 := createFood("Nasi Goreng", 12.000)
	food2 := createFood("Mie Ayam", 10.000)
	drink1 := createDrink("Kopi", 14.000)
	drink2 := createDrink("Teh", 8.000)

	info := []printer{food1, food2, drink1, drink2}

	for _, item := range info {
		item.printInfo()
	}
}
