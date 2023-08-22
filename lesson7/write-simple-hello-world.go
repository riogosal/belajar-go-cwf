package main

import "os"

func main() {
	stringKu := "hello"
	os.WriteFile("hello.txt", []byte(stringKu), 0644)
}
