package main

import "fmt"

func main() {
	//Con Channel Synchronization
	done := make(chan bool) // Creamos el canal para sincronizar
	fmt.Println("Inicia go routine del main")
	go hello(done)
	<-done //Espera hasta que hello() haya terminado
	fmt.Println("Termina go routine main")
}

func hello(done chan bool) {
	fmt.Println("Inicia go routine de Hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, "Hello World")
	}
	fmt.Println("Termina go routine de Hello")
	done <- true //Señala que ha terminado la go routine
}

func main2() {
	fmt.Println("Inicia Goroutine del main")
	go hello2()
	fmt.Println("Termina Goroutine del main")
}
func hello2() {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
}

/*
a) ¿Cuántas veces se imprime Hello world?
	Respuesta: 0

b) ¿Cuántas Goroutines tiene el programa?
	Respuesta: 2(main y hello)

c) ¿Cómo cambiaría el programa (con la misma cantidad de
Goroutines) para que imprima 3 veces Hello world?
i) Hágalo usando time.Sleep
	Respuesta: es igual al main2 y hello2 salvo que se le debe agregar la siguiente linea:
	go hello()
	time.Sleep(1 * time.Second) // Espera a que la Goroutine termine


ii) Hágalo usando Channel Synchronization
*/
