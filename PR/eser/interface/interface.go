package main

import "fmt"

type printer interface {
	printInfo() string
}

func Cetak(p printer) {
	fmt.Println("ini adalah :", p.printInfo())
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (d Drink) printInfo() string {
	return fmt.Sprintf("%s, %2f ", d.Name, d.Price)
}

func (b Book) printInfo() string {
	return fmt.Sprintf("\nBook  : %s Price : %.2f", b.Title, b.Price)
}
func main() {
	guslinger := Book{
		Title: "Eser Life Journey",
		Price: 12.000,
	}

	cofee := Drink{
		Name:  "Kopi Gula Aren",
		Price: 14.000,
	}

	Cetak(guslinger)
	Cetak(cofee)

	// guslinger.printInfo()
	// coffie.printInfo()

	// info := []printer{guslinger, coffie}
	// info[0].printInfo()
	// info[1].printInfo()
}
