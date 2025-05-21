package main

import "fmt"

func convert(v int, b int) string {
	if v < b {
		return fmt.Sprintf("%d", v) //El Sprintf convierte un valor al string de sus dígitos
		//El %d en Go indica que se debe imprimir un
		//argumento correspondiente a un entero en formato decimal
	}
	return convert(v/b, b) + fmt.Sprintf("%d", v%b)
}

func main() {
	var valor int
	var base int
	var s string
	fmt.Println("Ingrese valor: ")
	fmt.Scan(&valor)
	fmt.Println("Ingrese base (debe ser > a 1 y < a 37): ")
	fmt.Scan(&base)
	var signo string = "+"
	if valor < 0 {
		signo = "-"
		valor = -valor
	}
	s = signo + convert(valor, base)
	fmt.Println(s)
}
