package main

import (
	"fmt"
	"unicode"
)

func contarMayusculasYMinusculas(cadena string) (int, int) {
	runas := []rune(cadena)
	mayusculas := 0
	minusculas := 0
	for i, _ := range runas {
		if unicode.IsLetter(runas[i]) {
			if unicode.IsUpper(runas[i]) {
				mayusculas++
			} else {
				minusculas++
			}
		}
	}
	return mayusculas, minusculas
}

func main() {
	resultado := ""
	var letra string
	var ok bool
	for !ok {
		fmt.Println("Ingresar letra: ")
		fmt.Scan(&letra)
		if letra == "CD" {
			resultado += letra
			ok = true
		} else {
			resultado += letra
		}
	}
	letras := 0
	numeros := 0
	caracteresEspeciales := 0
	runas := []rune(resultado)
	digitos := make(map[rune]int)

	for i, _ := range runas {
		if unicode.IsLetter(runas[i]) {
			letras++
		} else if unicode.IsNumber(runas[i]) {
			numeros++
			digitos[runas[i]]++
		} else {
			caracteresEspeciales++
		}
	}
	fmt.Println("Cantidad de numeros: ", numeros)
	fmt.Println("Cantidad de letras: ", letras)
	fmt.Println("Cantidad de caracteres especiales: ", caracteresEspeciales)
	var mayusculas int
	var minusculas int
	mayusculas, minusculas = contarMayusculasYMinusculas(resultado)
	fmt.Println("Cantidad de mayusculas: ", mayusculas)
	fmt.Println("Cantidad de minusculas: ", minusculas)
	fmt.Println("Ocurrencias de cada digito: ")
	for i, cant := range digitos {
		fmt.Printf("Dígito %c: %d veces\n", i, cant)
	}
}
