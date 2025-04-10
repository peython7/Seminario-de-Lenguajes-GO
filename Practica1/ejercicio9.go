package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var frase string
	fmt.Println("Ingrese una frase:")
	// Lee la frase completa
	fmt.Scanln(&frase)

	// Si querés que tome toda la línea (incluyendo espacios), podés usar:
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// frase = scanner.Text()

	resultado := reemplazarJueves(frase)

	fmt.Println("Frase modificada:")
	fmt.Println(resultado)
}

func reemplazarJueves(frase string) string {
	buscada := "jueves"
	reemplazo := "martes"

	var resultado strings.Builder
	i := 0

	for i < len(frase) {
		// Si encuentra "jueves" (sin importar mayúsculas/minúsculas)
		if len(frase)-i >= len(buscada) && strings.EqualFold(frase[i:i+len(buscada)], buscada) {
			// Construye la palabra reemplazo respetando las mayúsculas/minúsculas de la original
			for j := 0; j < len(reemplazo); j++ {
				if unicode.IsUpper(rune(frase[i+j])) {
					resultado.WriteByte(byte(unicode.ToUpper(rune(reemplazo[j]))))
				} else {
					resultado.WriteByte(reemplazo[j])
				}
			}
			i += len(buscada)
		} else {
			// Sino, sigue copiando la frase original
			resultado.WriteByte(frase[i])
			i++
		}
	}

	return resultado.String()
}
