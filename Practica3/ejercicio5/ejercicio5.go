package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func productor(id int, ch chan<- int, wg *sync.WaitGroup) {
	// Cada vez que una goroutine termina, llama a wg.Done(), que resta 1 al contador.
	// El wg.Done, cuando se ejecuta, le resta uno al contador de wg de goroutines activas
	defer wg.Done() // el defer pospone la ejecución de la función wg.Done hasta que la función que lo contiene termina.
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // duerme entre 0 y 1 segundo
		num := rand.Intn(101)                                         // Genera un número entre 0 y 100
		fmt.Printf("Productor %d produjo: %d\n", id, num)
		ch <- num //Manda el número al canal (ch <- num).
	}
}

func consumidor(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		num := <-ch // recibe un número
		fmt.Printf("Consumidor %d consumió: %d\n", id, num)
	}
}

func main() {
	ch := make(chan int, 2) // canal con buffer pequeño (puede ser 2 o más).
	var wg sync.WaitGroup   //Sirve para esperar que un conjunto de goroutines terminen antes de seguir.

	wg.Add(4) //Suma 4 al contador interno del WaitGroup.
	go productor(1, ch, &wg)
	go productor(2, ch, &wg)
	go consumidor(1, ch, &wg)
	go consumidor(2, ch, &wg)

	wg.Wait() // espera que terminen productores y consumidores
	fmt.Println("Todos los productores y consumidores han terminado.")
}

/*EXPLICACION DE time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)!!!
rand.Intn(1000):
	Genera un número aleatorio entre 0 y 999 inclusive.

time.Duration(rand.Intn(1000)):
	Convierte ese número en un tipo duración (Duration es solo un int64 detrás de escena).

time.Duration(rand.Intn(1000)) * time.Millisecond:
	Convierte esos milisegundos en una duración en milisegundos reales.

time.Sleep(...)
	Pone a dormir la goroutine por esa cantidad de tiempo.
*/
