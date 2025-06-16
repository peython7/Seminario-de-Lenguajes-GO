package main

import (
	"fmt"
	"time"
)

// función que envía una secuencia de números enteros por canal
func enviarSecuencia(ch chan int, valores []int) {
	for _, v := range valores {
		ch <- v
		time.Sleep(300 * time.Millisecond) // simula trabajo
	}
	close(ch) // se avisa que no se envían más datos
}

func main() {
	// Declaración de los canales
	canal1 := make(chan int)
	canal2 := make(chan int)
	canal3 := make(chan int)

	// Iniciamos las goroutines productoras
	go enviarSecuencia(canal1, []int{1, 2, 9, 8})
	go enviarSecuencia(canal2, []int{7, 9, 3})
	go enviarSecuencia(canal3, []int{100, 200, 300})

	contador1 := 0
	contador2 := 0
	contador3 := 0

	// Control de canales abiertos
	abierto1 := true
	abierto2 := true
	abierto3 := true

	for abierto1 || abierto2 || abierto3 {
		select {
		case v, ok := <-canal1:
			if ok {
				fmt.Println("Recibido del canal 1:", v)
				contador1++
			} else {
				abierto1 = false
			}
		case v, ok := <-canal2:
			if ok {
				fmt.Println("Recibido del canal 2:", v)
				contador2++
			} else {
				abierto2 = false
			}
		case v, ok := <-canal3:
			if ok {
				fmt.Println("Recibido del canal 3:", v)
				contador3++
			} else {
				abierto3 = false
			}
		}
	}
	fmt.Println("\nResumen final:")
	fmt.Println("Total recibido del canal 1:", contador1)
	fmt.Println("Total recibido del canal 2:", contador2)
	fmt.Println("Total recibido del canal 3:", contador3)
}
