package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name     string
	Age      int
	Username string
	Password string
}

func main() {
	peopleBytes, err := os.ReadFile("people.json")
	if err != nil {
		panic(err)
	}

	var people []Person
	if err := json.Unmarshal(peopleBytes, &people); err != nil {
		panic(err)
	}

	found := false
	for _, p := range people {
		if p.Name == "Eser" {
			found = true
		}
	}
	if found {
		fmt.Println("Eser found")
	} else {
		fmt.Println("Eser not found")
	}

}
