package main

import "fmt"

type Map[K comparable, V any] map[K]V

func main() {
	//Primer uso
	var edades Map[string, int] = Map[string, int]{
		"Pedro":  24,
		"Carlos": 51,
	}
	fmt.Println("Edades: ")

	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)
	}

	//Segundo uso
	var nombres Map[int, string] = Map[int, string]{
		43044473: "Pedro Vega",
		23583811: "Carlos Vega",
	}
	fmt.Println("Nombres: ")
	for dni, nombre := range nombres {
		fmt.Printf("La persona con dni %d se llama %s\n", dni, nombre)
	}
}
