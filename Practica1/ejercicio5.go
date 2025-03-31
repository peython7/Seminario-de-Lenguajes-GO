package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Println("Ingrese un numero :")
	fmt.Scan(&num)

	// switch para seleccionar la condición
	switch {
	case num == 0:
		fmt.Println("El número es 0")
		return
	case num < -18:
		num = math.Abs(num) // valor absoluto
	case num >= -18 && num <= -1:
		num = math.Mod(num, 4) // modulo con flotante
	case num >= 1 && num < 20:
		num = math.Pow(num, 2) // cuadrado
	default:
		num = -num
	}

	fmt.Println("El numero resultante es:", num)
}
