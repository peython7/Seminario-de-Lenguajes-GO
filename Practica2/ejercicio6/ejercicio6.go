package main

import "fmt"

func Sum(v1, v2 []int) []int {
	var min int
	if len(v1) < len(v2) {
		min = len(v1)
	} else {
		min = len(v2)
	}
	resultado := make([]int, min)

	for i := 0; i < min; i++ {
		resultado[i] = v1[i] + v2[i]
	}
	return resultado
}

func Avg(v []int) int {
	var resultado int
	for i := 0; i < len(v); i++ {
		resultado += v[i]
	}
	return resultado / len(v)
}

func AvgFloat(v []int) float64 {
	var resultado float64
	for i := 0; i < len(v); i++ {
		resultado += float64(v[i])
	}
	return resultado / float64(len(v))
}

func main() {
	slice1 := []int{4, 5, 7, 8, 9}
	slice2 := []int{3, 4, 9, 7}
	resultado := Sum(slice1, slice2)
	fmt.Println(resultado)
	fmt.Println("El promedio del slice ", slice1, " es de: ", Avg(slice1))
	fmt.Println("El promedio del slice ", slice1, " en float64 es de: ", AvgFloat(slice1))
}
