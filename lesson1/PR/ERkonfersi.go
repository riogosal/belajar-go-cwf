// Rumus C = 9/5 * (f-32)
package main

import (
	"fmt"
)

func main() {
	var F float32 = 30
	var C float32 = (float32(F) - 32) * 5 / 9

	fmt.Printf("%.f Fahrenheit = %.2f Celsius", F, C)
	// f adalah nilai yang digantikan adalah float atau bilangan pecahan.
	// 2 angka desimal di belakang koma
	// '%' = tanda bahwa di situ ada suatu nilai yang akan dimasukkan dan menggantikan tanda % tersebut saat mencetak.
}
