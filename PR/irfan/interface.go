package main

import "fmt"

type HasName interface {
	GetName() string
}

func HaloGuys(h HasName) {
	fmt.Println("Hello ", h.GetName())
}

type User struct {
	Name string
}

func (u User) GetName() string {
	return u.Name
}

func main() {
	var angger User
	angger.Name = "Angger"

	HaloGuys(angger)

	eser := User{
		Name: "eser",
	}

	HaloGuys(eser)
	HaloGuys(angger)
}
