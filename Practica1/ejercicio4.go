package main

import "fmt"

const LIMITE = 250

func main() {
	suma := 0

	for i := 2; i <= LIMITE; i += 2 {
		suma += i
	}
	fmt.Println("Suma de los primeros pares (ascendente):", suma)

	sumaInversa := 0
	for i := LIMITE; i >= 2; i -= 2 {
		sumaInversa += i
	}
	fmt.Println("Suma de los primeros pares (descendente):", sumaInversa)
}
