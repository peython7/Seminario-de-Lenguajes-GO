package main

import (
	"fmt"
)

type (
	Celsius    float64
	Fahrenheit float64
)

func agregarTemperaturas(temps *[10]Celsius) {
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese una temperatura: ")
		fmt.Scan(temps[i])
	}
}

func CelsiusToFahrenheit(t Celsius) {
	f := (t * 9 / 5) + 32
	fmt.Println("La temperatura ", t, " en fahrenheit es: ", f)
}

func main() {
	const totalPacientes = 10
	var temperaturas [totalPacientes]Celsius
	agregarTemperaturas(&temperaturas)

	var altas, normales, bajas int
	var maxT, minT Celsius = 0, 1000

	for i := 0; i < totalPacientes; i++ {
		if temperaturas[i] < 36 {
			bajas++
		} else if temperaturas[i] > 37.5 {
			altas++
		} else {
			normales++
		}
		if maxT < temperaturas[i] {
			maxT = temperaturas[i]
		}
		if minT > temperaturas[i] {
			minT = temperaturas[i]
		}
	}

	for i := 0; i < totalPacientes; i++ {
		CelsiusToFahrenheit(temperaturas[i])
	}

	fmt.Println(altas)
	fmt.Println(bajas)
	fmt.Println(normales)
	fmt.Println(maxT, " y ", minT)

	fmt.Println("El porcentaje de temperaturas altas es de :", (altas*100)/totalPacientes)
	fmt.Println("El porcentaje de temperaturas bajas es de :", (bajas*100)/totalPacientes)
	fmt.Println("El porcentaje de temperaturas normales es de :", (normales*100)/totalPacientes)

	prom := int(((maxT - minT) / maxT) * 100)
	fmt.Print("El porcentaje entre las temperaturas maxima y minima es de: ", prom)
}
