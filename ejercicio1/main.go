package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto1 string
var status1 bool = true

func main() {
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false

	var moneda int64 = 21

	numero2 = int(moneda)

	texto = fmt.Sprintf("%d", moneda)

	texto = strconv.Itoa(numero2)

	fmt.Println(numero2)
	fmt.Println(texto1)
	fmt.Println(status1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
}

func MostrarStatus() {
	fmt.Println(status1)
}
