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

func menu(p Print) {
	p.printInfo()
}

func (f Food) printInfo() {
	fmt.Printf("Food  : %s Price : %.2f\n", f.Name, f.Price)
}

func (d Drink) printInfo() {
	fmt.Printf("Drink : %s  Price : %.2f\n", d.Name, d.Price)
}

// factory faction
func MyFood(name string, price float32) Print {
	return Food{Name: name, Price: price}
}

func MyDrink(name string, price float32) Print {
	return Drink{Name: name, Price: price}
}

func main() {
	food1 := MyFood("Nasi Goreng", 12.000)
	drink1 := MyDrink("Kopi", 14.000)
	// food2 := MyFood("Mie Ayam", 10.000)
	// drink2 := MyDrink("Teh", 8.000)

	menu(food1)
	menu(drink1)
	// menu(food2)
	// menu(drink2)

	// info := []Print{food1, food2, drink1, drink2}

	// for _, item := range info {
	// 	item.printInfo()
	// }
}
