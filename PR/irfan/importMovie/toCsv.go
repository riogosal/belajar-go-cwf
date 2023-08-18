package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name, Username, Password string
	Age                      int
}

func (p Person) ToCSV() string {
	return fmt.Sprintf("%s,%d,%s,%s\n", p.Name, p.Age, p.Username, p.Password)
}

func FromCSV(csv string) Person {
	var p Person
	fmt.Sscanf(csv, "%s,%d,%s,%s\n", &p.Name, &p.Age, &p.Username, &p.Password)
	return p
}

func main() {

	var dataName Person

	for true {
		fmt.Println("1. create data, 2. exit :")
		var choice string
		fmt.Scanln(&choice)
		if choice == "1" {
			fmt.Println("masukkan nama :")
			var name string
			fmt.Scanln(&name)
			dataName.Name = name

			fmt.Println("masukkan Age :")
			var age int
			fmt.Scanln(&age)
			dataName.Age = age

			var username string
			fmt.Println("masukkan Username :")
			fmt.Scanln(&username)
			dataName.Username = username

			fmt.Println("masukkan Password :")
			var password string
			fmt.Scanln(&password)
			dataName.Password = password
		} else {
			break
		}

	}
	// import data from csv

	arrayDataPerson := []Person{dataName}

	f, err := os.Create("dataFilm.csv")

	if err != nil {
		panic("error " + err.Error())

	}

	for _, p := range arrayDataPerson {

		f.WriteString(p.ToCSV())
	}
	fmt.Println("Program successfully write people.csv")
}
