package main

import (
	"fmt"
	"strconv"
)

func mainInteger() {
	x := 10
	y := 20.0
	z := x + int(y)

	println(z)
}

func mainFloat() {
	x := 10
	y := 20.0
	z := float64(x) + y

	println(z)
}

func mainChar() {
	x := 'a'
	y := 'b'
	z := x + y

	println(z)
}

func mainString() error {
	x := "abc"
	y := "def"
	z := 123
	gabunganKata := x + y + fmt.Sprint(z)
	test := x[0]
	println(test)
	println(string(test))

	println(gabunganKata)

	bilangan, err := strconv.Atoi("123a")
	if err != nil {
		// sama dengan die() atau exit() di PHP
		return err
	}
	println("Sampe ma disini")
	println(bilangan + 123) // 123 + 123 = 246

	return nil
}

func main() {
	// mainInteger()
	// mainFloat()
	mainChar()
	// Jangan hiraukan error seperti line di bawah ini
	mainString()
	// Ini yang betul
	if err := mainString(); err != nil {
		println("Woi ada error we")
		panic(err)
	}

	println("Santai")
}
