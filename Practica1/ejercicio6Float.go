package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	fmt.Println(num1, "", num2)
	if num1 < num2 {
		fmt.Println(num2 / num1)
	} else {
		fmt.Println(num1 / num2)
	}
}
