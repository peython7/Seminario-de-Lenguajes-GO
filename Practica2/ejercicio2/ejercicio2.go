package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * (factorial(n - 1))
	}
}

func main() {
	var n int
	fmt.Println("Ingresar numero: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("ERROR")
	} else if n == 0 || n == 1 {
		fmt.Println(n)
	} else {
		resultado := 1
		for i := 2; i <= n; i++ {
			resultado *= i
		}
		fmt.Println("El factorial de ", n, " es: ", resultado)
	}
	fmt.Println("El factorial de", n, "es: ", factorial(n))
}
