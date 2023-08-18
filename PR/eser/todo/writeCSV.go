package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name     string
	Age      int
	Username string
	Password string
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
	johnDoe := Person{
		Name:     "Eser",
		Age:      28,
		Username: "eser",
		Password: "password",
	}
	people := []Person{
		johnDoe,
		{
			Name:     "Jane Doe",
			Age:      20,
			Username: "janedoe",
			Password: "password",
		},
	}

	file := "myData.csv"

	f, err := os.Create(file)

	if err != nil {
		fmt.Println("Gagal membuat File:", err)
		return
	}

	defer f.Close()

	for _, person := range people {
		_, err := f.WriteString(person.ToCSV())
		if err != nil {
			fmt.Println("Gagal menulis data ke dalam file :", err)
			return
		}
	}

	fmt.Println("Program successfully write people.csv")
}
