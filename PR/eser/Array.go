package main

import "fmt"

func main() {
	// myArray()
	// myFruit()
	myMap()
	// myStruct()
}

func myArray() {
	numbers := []string{
		"apel", "mangga", "jeruk", "pir",
	}
	fmt.Println(len(numbers))
	for _, value := range numbers {
		if value != "" {
			fmt.Println(value)
			// fmt.Printf("Index: %d, Value: %d\n", i, value)
		}
	}
	// var numbers [5]int
	// numbers[0] = 1
	// numbers[1] = 2
	// numbers[2] = 3
	// numbers[3] = 4
	// numbers[4] = 5

	// fmt.Println(numbers[3])
}

func myMap() {
	// myList := map[string]string{
	// 	"Eser":   "Human",
	// 	"Jarvis": "Ai",
	// }

	// fmt.Println(myList["jarvis"])

	person := map[string]string{
		"Name":    "Eser",
		"Address": "Bali",
	}

	// person["Address"] = "Toraja"
	// person["Name"] = "Eliaser Randa"

	if person["Address"] == "Bali" {
		// delete(person, "Address")
		person["Address"] = "Unknow"
	}
	// fmt.Println(person)
	fmt.Println(person["Name"])
	fmt.Println(person["Address"])
}

type Person struct {
	Name string
	age  int
}

type Fruit struct {
	Name   string
	Amount int
}

func myFruit() {
	fruit := Fruit{
		Name:   "Apel",
		Amount: 100,
	}

	fmt.Println(fruit.Name, fruit.Amount)
}

func myStruct() {
	person := Person{
		Name: "Igor",
		age:  17,
	}
	// fmt.Println(person.Name, person.age)
	fmt.Printf("Nama saya adalah %s Umur %d", person.Name, person.age)
}

//kebunBinatang := []Animal{}
