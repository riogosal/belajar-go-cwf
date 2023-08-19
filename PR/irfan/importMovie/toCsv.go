package main

import (
	"encoding/json"
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
	arrayDataPerson := []Person{}

	for {
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

			arrayDataPerson = append(arrayDataPerson, dataName)
		} else {
			break
		}

	}
	// import data from csv

	f, err := os.Create("dataFilm.csv")

	if err != nil {
		panic("error " + err.Error())

	}

	for _, p := range arrayDataPerson {

		f.WriteString(p.ToCSV())

		fmt.Println(FromCSV(p.ToCSV()))
	}
	fmt.Println("Program successfully write people.csv")

	fileJson, err := json.Marshal(arrayDataPerson)

	if err != nil {
		panic("error " + err.Error())

	}
	os.WriteFile("fileJson.json", fileJson, 0644)
}
