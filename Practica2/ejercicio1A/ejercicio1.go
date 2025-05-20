package main

import (
	"fmt"
)

func main() {
	const totalPacientes = 10
	var t float64
	var pacientesPorTemperatura [4]int
	for i := 0; i < totalPacientes; i++ {
		fmt.Println("Ingresar temperatura: ")
		fmt.Scan(t)
		if t > 50 || t < 20 {
			pacientesPorTemperatura[3]++
		} else if t < 36 {
			pacientesPorTemperatura[0]++
		} else if t > 37.5 {
			pacientesPorTemperatura[1]++
		} else if t < 37.5 && t > 36 {
			pacientesPorTemperatura[2]++
		}
	}
}
