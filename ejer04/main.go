package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var leyenda string
var resultado int

func main() {
	fmt.Println("ingrese el num 1")
	fmt.Scanf("%d", &numero1)

	fmt.Println("ingrese el num 2")
	fmt.Scanf("%d", &numero2)

	fmt.Println("Descripcion: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		leyenda = scanner.Text()
		fmt.Printf("%s", leyenda)
	}
	resultado = numero1 + numero2
	fmt.Printf("%d", numero1+numero2)

}
