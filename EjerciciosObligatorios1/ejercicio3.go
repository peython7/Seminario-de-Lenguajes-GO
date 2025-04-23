package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func modificarPalabra(palabra string) string {
	runas := []rune(palabra)
	var resultado strings.Builder
	for i, _ := range runas { // Cambiar
		if unicode.IsUpper(runas[i]) {
			resultado.WriteRune(unicode.ToLower(runas[i]))
		} else {
			resultado.WriteRune(unicode.ToUpper(runas[i]))
		}
	}
	return resultado.String()
}

func limpiarPalabra(palabra string) (string, string) {
	runas := []rune(palabra)
	var puntuacion strings.Builder
	var resultado strings.Builder
	for i, _ := range runas {
		if !unicode.IsPunct(runas[i]) {
			resultado.WriteRune(runas[i])
		} else {
			puntuacion.WriteRune(runas[i])
		}
	}
	return resultado.String(), puntuacion.String()
}

func main() {
	argumento := os.Args[1]
	frase := "Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO"
	palabras := strings.Fields(frase)
	var resultado strings.Builder
	for i := 0; i < len(palabras); i++ {
		palabraLimpia, puntuacion := limpiarPalabra(palabras[i])
		if strings.EqualFold(argumento, palabraLimpia) {
			resultado.WriteString(modificarPalabra(palabraLimpia))
			if puntuacion != "" {
				resultado.WriteString(puntuacion + " ")
			} else {
				resultado.WriteString(" ")
			}
		} else {
			resultado.WriteString(palabras[i] + " ")
		}
	}
	fmt.Print(resultado.String())
}
