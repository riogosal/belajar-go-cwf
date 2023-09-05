package main

import (
	"fmt"
	"os"
)

type ExportableToExcel interface {
	ExportToExcel() string
}

type Printable interface {
	Print() string
}

type Data struct {
	NoIkan   int
	Grade    string
	NoBasket int
}

type BluetoothPrinter struct {
	Buffer []byte
}

func (bp BluetoothPrinter) Print() string {
	return fmt.Sprintf(`
		<html>
			<content>Ikan No: %d</content>
			<content>Grade: %s</content>
			<content>No. Basket: %d</content>
		</xml>
	`, bp.Buffer[0], bp.Buffer[1], bp.Buffer[2])
}

type ThermalPrinter struct {
	Data Data
}

func (tp ThermalPrinter) Print() string {
	return fmt.Sprintf(`
		<xml>
			<content>Ikan No: %d</content>
			<content>Grade: %s</content>
			<content>No. Basket: %d</content>
		</xml>
	`, tp.Data.NoIkan, tp.Data.Grade, tp.Data.NoBasket)
}

type LabelPrinter struct {
	Message Data
}

func (lp LabelPrinter) Print() string {
	return fmt.Sprintf(`^XA
	^FO50,50
	^FDIkan No: %d^FS
	^FO50,100
	^FDGrade: %s^FS
	^FO50,150
	^FDNo. Basket: %d^FS
	^XZ
	`, lp.Message.NoIkan, lp.Message.Grade, lp.Message.NoBasket)
}

func PrintThermalPrinter(tp ThermalPrinter) {
	fmt.Println(tp.Print())
}
func PrintLabelPrinter(lp LabelPrinter) {
	fmt.Println(lp.Print())
}

func PrintToPrinter(printer Printable) {
	file, err := os.OpenFile("./printer", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	file.WriteString(printer.Print())
}

func main() {
	labelPrinter := LabelPrinter{
		Message: Data{
			NoIkan:   1,
			Grade:    "A",
			NoBasket: 1,
		},
	}

	PrintToPrinter(labelPrinter)

	thermalPrinter := ThermalPrinter{
		Data: Data{
			NoIkan:   5,
			Grade:    "B",
			NoBasket: 13,
		},
	}
	PrintToPrinter(thermalPrinter)
}
