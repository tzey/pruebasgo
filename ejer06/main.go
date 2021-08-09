package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("la suma de 3 y 5 es %d \n", Calculo(3, 5))

	fmt.Printf("la suma de 3 y 5 es %d \n", Suma(3, 5))
}

func Suma(num1 int, num2 int) int {
	return num1 + num2
}
