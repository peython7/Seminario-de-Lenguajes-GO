package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	frase := "El miércoLes me gusta porque es Miércoles"
	reemplazo := "automóvil"
	runesReemplazo := []rune(reemplazo)
	buscado := "miércoles"
	var resultado strings.Builder
	for {
		pos := strings.Index(strings.ToLower(frase), buscado)
		if pos != -1 { //si da -1 es porque no encontro otra ocurrencia
			resultado.WriteString(frase[:pos])
			aux := []rune(frase[pos : pos+len(buscado)])
			for i, _ := range aux { // Cambiar
				if unicode.IsUpper(aux[i]) {
					resultado.WriteRune(unicode.ToUpper(runesReemplazo[i]))
				} else {
					resultado.WriteRune(unicode.ToLower(runesReemplazo[i]))
				}
			}
			frase = frase[pos+len(buscado):]
		} else {
			resultado.WriteString(frase) // agrego lo que falta
			break
		}
	}
	fmt.Print(resultado.String())
}
