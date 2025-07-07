package main

import (
	"fmt"
	"strings"
)

func invertirPalabra(palabra string) string {
	runas := []rune(palabra)
	j := len(runas) - 1
	i := 0
	ok := false
	for ok == false {
		if i > j {
			ok = true
		} else {
			runas[i], runas[j] = runas[j], runas[i]
			i++
			j--
		}
	}
	return string(runas)
}

func main() {
	frase := "Qué lindo día es hoy"
	palabras := strings.Fields(frase) //el strings.Fields divide la cadena en un slice de palabras, separándolas por espacios en blanco 
	var resultado strings.Builder
	for i := 0; i < len(palabras); i++ {
		if (i+1)%2 != 0 {
			resultado.WriteString(invertirPalabra(palabras[i]) + " ")
		} else {
			resultado.WriteString(palabras[i] + " ")
		}
	}
	fmt.Print(resultado.String())
}
