package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var resultado int
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	fmt.Println(num1, "", num2)
	if num1 == 0 && num2 > 0 {
		fmt.Println("No se puede realizar una division por 0")
	} else if num2 == 0 && num1 > 0 {
		fmt.Println("No se puede realizar una division por 0")
	} else if num1 < num2 {
		resultado = num2 / num1
		fmt.Println(resultado)
	} else {
		resultado = num1 / num2
		fmt.Println(resultado)
	}
}
