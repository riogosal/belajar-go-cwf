package main

import "os"

func WriteSimpleHelloWorld() {
	stringKu := "hello\nworld"
	os.WriteFile("hello.txt", []byte(stringKu), 0644)
}
