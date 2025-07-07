package main

import (
	"fmt"     // para leer argumentos desde la línea de comandos
	"strconv" // para convertir texto a número
	"time"    // para medir el tiempo de ejecución
)

func esPrimo(n int) bool {
	if n < 2 { // los numeros < 2 no son primos
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now() // iniciamos el cronómetro para medir el tiempo de ejecución
	var input string    // aca almacenamos lo que vamos a leer
	fmt.Print("Ingresá un número entero positivo: ")
	fmt.Scanln(&input) // Ponemos &input porque Scanln necesita la dirección de memoria donde va a guardar lo que lea

	N, err := strconv.Atoi(input) // Convertimos ese string (por ejemplo, "20") a un número entero (int).

	if err != nil || N < 1 { // si se lee por ejemplo un string en err no se guarda nil
		fmt.Println("Entrada inválida. Debe ser un entero positivo.")
		return
	}

	fmt.Printf("Números primos menores o iguales a %d:\n", N) // calculamos los primos
	//%d es un placeholder para un número entero.
	//\n agrega salto de línea.
	for i := 2; i <= N; i++ {
		if esPrimo(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	elapsed := time.Since(start)
	fmt.Println("Tiempo total:", elapsed)
}
