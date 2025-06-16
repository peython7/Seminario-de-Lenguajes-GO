package main

import (
	"fmt"
	"math/rand"
	"time"
)

func enviarDatos(ch chan string, nombre string, intervalo time.Duration) {
	for {
		ch <- fmt.Sprintf("%s -> %d", nombre, rand.Intn(1000))
		time.Sleep(intervalo)
	}
}

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go enviarDatos(canal1, "Canal 1", 500*time.Millisecond) // cada 0.5s
	go enviarDatos(canal2, "Canal 2", 1*time.Second)        // cada 1s
	//La diferencia de tiempos es para que se vea mas fluido al momento de ejecutar

	// Definimos los timers de 5s y 10s
	timeout1 := time.After(5 * time.Second)
	timeout2 := time.After(10 * time.Second)

	for {
		select {
		case msg := <-canal1:
			fmt.Println("Recibido de ch1:", msg)
		case msg := <-canal2:
			fmt.Println("Recibido de ch2:", msg)
		case <-timeout1:
			fmt.Println("Tiempo finalizado para canal 1 (5s)")
			canal1 = nil // anulamos el canal 1 para que no se seleccione más
		case <-timeout2:
			fmt.Println("Tiempo finalizado para canal 2 (10s)")
			return // salimos del programa
		}
	}
}
