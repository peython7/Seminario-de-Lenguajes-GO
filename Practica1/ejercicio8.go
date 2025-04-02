package main

import (
	"fmt"
)

func main() {
	var viento string
	fmt.Print("Ingrese la dirección del viento (N, S, E, O): ")
	fmt.Scanln(&viento)

	switch viento {
	case "N":
		fmt.Println("El viento va hacia el Norte")
	case "S":
		fmt.Println("El viento va hacia el Sur")
	case "E":
		fmt.Println("El viento va hacia el Este")
	case "O":
		fmt.Println("El viento va hacia el Oeste")
	default:
		fmt.Println("Dirección no válida")
	}
}
