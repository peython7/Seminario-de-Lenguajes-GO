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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Debe ingresar una palabra como argumento.")
		return
	}
	argumento := os.Args[1]
	frase := "Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO"
	var resultado strings.Builder
	for {
		pos := strings.Index(strings.ToLower(frase), strings.ToLower(argumento))
		if pos != -1 { //si da -1 es porque no encontro otra ocurrencia
			resultado.WriteString(frase[:pos])
			palabraRuna := (frase[pos : pos+len(argumento)])
			resultado.WriteString(modificarPalabra(palabraRuna))
			frase = frase[pos+len(argumento):]
		} else {
			resultado.WriteString(frase) // agrego lo que falta
			break
		}
	}
	fmt.Print(resultado.String())
}

/*
os.Args es un slice de strings con los argumentos de línea de comandos.

os.Args[0] es el nombre del programa.

os.Args[1] es el primer argumento real.

Entonces, si len(os.Args) < 2, significa que no se pasó ningún argumento, y el programa muestra un mensaje y termina.
*/
