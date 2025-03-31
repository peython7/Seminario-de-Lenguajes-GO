package main

import "fmt"

func main() {
	var temperaturas [3]float64
	var i int
	var max float64 = 0
	var min float64 = 1000
	var temp float64
	for i = 0; i < 10; i++ {
		fmt.Print("Ingrese temperatura del paciente: ")
		fmt.Scan(&temp)
		switch {
		case temp > 37.5:
			temperaturas[0]++
		case temp <= 37.5 && temp > 36:
			temperaturas[1]++
		case temp <= 36:
			temperaturas[2]++
		}
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}
	fmt.Println("Porcentaje del grupo +37.5: ", float64(temperaturas[0]/10))
	fmt.Println("Porcentaje del grupo +36: ", float64(temperaturas[1]/10))
	fmt.Println("Porcentaje del grupo -36: ", float64(temperaturas[2]/10))
	fmt.Println("Maximo: ", max, " Minimo: ", min)
	fmt.Println("Porcentaje entro maximo y minimo: ", int(max/min))
}
