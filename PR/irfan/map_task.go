package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name map[string]string, filter Filter) {
	for _, v := range name {
		fmt.Println("Hello My Name is : ", filter(v))
	}

}

func spamFilte(name string) string {

	if name == "goblok" {
		return "..."
	} else {
		return strings.ToUpper(name)
	}

}

func main() {

	listName := map[string]string{
		"name":  "irfan",
		"name2": "angger",
		"name3": "adam",
		"name4": "goblok",
	}

	sayHelloWithFilter(listName, spamFilte)

}
