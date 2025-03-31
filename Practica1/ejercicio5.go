package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Println("Ingrese un numero :")
	fmt.Scan(&num)
	//var infinito float64 = math.Inf(1)  INFINITO
	//var menosInfinito float64= math.Inf(-1)  INFINITO NEGATIVO
	if num == 0 {
		fmt.Println("El numero es 0")
		return
	} else if num < -18 {
		num = math.Abs(num) //valor absoluto
	} else if (num >= -18) && (num <= -1) {
		num = float64(math.Mod(num, 4)) //a num le doy valor de entero y luego al resultado float64
	} else if (num >= 1) && (num < 20) {
		num = math.Pow(num, 2) //Incrementa al cuadrado
	} else {
		num = -num
	}
	fmt.Println("El numero resultante es: ", num)
}
