package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	frase := "Hoy es Jueves y en siete dias sera jueVes tambien."
	reemplazo := "martes"
	buscado := "jueves"
	var resultado strings.Builder
	for {
		pos := strings.Index(strings.ToLower(frase), buscado)
		if pos != -1 { //si da -1 es porque no encontro otra ocurrencia
			resultado.WriteString(frase[:pos])
			for i := 0; i < len(buscado); i++ {
				if unicode.IsUpper(rune(frase[pos+i])) {
					resultado.WriteRune(unicode.ToUpper(rune(reemplazo[i])))
				} else {
					resultado.WriteRune(unicode.ToLower(rune(reemplazo[i])))
				}
			}
			frase = frase[pos+len(buscado):]
		} else {
			resultado.WriteString(frase) // agrego lo que falta
			break
		}
	}
	for j := 0; j < len(resultado.String()); j++ {
		fmt.Print(string(resultado.String()[j]))
	}
}
