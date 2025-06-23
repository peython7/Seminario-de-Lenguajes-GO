package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Tarea struct {
	Numero int
}

var (
	mutexArchivo0  sync.Mutex
	mutexArchivo1  sync.Mutex
	mutexAcumulado sync.Mutex
	acumulado      int
	wg             sync.WaitGroup
)

//FUNCIONES POR PRIORIDAD

func procesarCanal0(ch <-chan Tarea, f *os.File) {
	for t := range ch {
		resultado := sumarDigitos(t.Numero)
		mutexArchivo0.Lock()
		f.WriteString(fmt.Sprintf("(0, %d)\n", resultado))
		mutexArchivo0.Unlock()
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}
}

func procesarCanal1(ch <-chan Tarea, f *os.File) {
	for t := range ch {
		resultado := invertirDigitos(t.Numero)
		mutexArchivo1.Lock()
		f.WriteString(fmt.Sprintf("(1, %d)\n", resultado))
		mutexArchivo1.Unlock()
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}
}

func procesarCanal2(ch <-chan Tarea) {
	for t := range ch {
		fmt.Printf("Prioridad 2: %d * 10 = %d\n", t.Numero, t.Numero*10)
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}
}

func procesarCanal3(ch <-chan Tarea) {
	for t := range ch {
		mutexAcumulado.Lock()
		acumulado += t.Numero
		fmt.Printf("Prioridad 3 acumulado (+%d): %d\n", t.Numero, acumulado)
		mutexAcumulado.Unlock()
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}
}

func sumarDigitos(n int) int {
	suma := 0
	for n > 0 {
		suma += n % 10
		n /= 10
	}
	return suma
}

func invertirDigitos(n int) int {
	invertido := 0
	for n > 0 {
		invertido = invertido*10 + n%10
		n /= 10
	}
	return invertido
}

func main() {

	file0, _ := os.Create("prioridad0.txt")
	file1, _ := os.Create("prioridad1.txt")
	defer file0.Close()
	defer file1.Close()

	// Creamos los canales por prioridad
	canal0 := make(chan Tarea)
	canal1 := make(chan Tarea)
	canal2 := make(chan Tarea)
	canal3 := make(chan Tarea)

	// Iniciamos procesamiento paralelo por canal
	go procesarCanal0(canal0, file0)
	go procesarCanal1(canal1, file1)
	go procesarCanal2(canal2)
	go procesarCanal3(canal3)

	// Creamos las listas por prioridad
	var tareas0, tareas1, tareas2, tareas3 []Tarea

	cantTareas := rand.Intn(81) + 20
	fmt.Println("Cantidad total de tareas generadas:", cantTareas)

	for i := 0; i < cantTareas; i++ {
		n := rand.Intn(1000)
		p := rand.Intn(4)
		t := Tarea{Numero: n}

		fmt.Printf("Tarea Prioridad: %d, Valor: %d\n", p, n)
		time.Sleep(100 * time.Millisecond)

		// Clasificamos por prioridad
		switch p {
		case 0:
			tareas0 = append(tareas0, t)
		case 1:
			tareas1 = append(tareas1, t)
		case 2:
			tareas2 = append(tareas2, t)
		case 3:
			tareas3 = append(tareas3, t)
		}
	}

	// Enviamos en orden:
	// prioridad 0
	fmt.Println("Procesando tareas de prioridad 0...")
	for _, t := range tareas0 {
		wg.Add(1)
		canal0 <- t
	}
	wg.Wait()

	//prioridad 1
	fmt.Println("Procesando tareas de prioridad 1...")
	for _, t := range tareas1 {
		wg.Add(1)
		canal1 <- t
	}
	wg.Wait()

	//prioridad 2
	fmt.Println("Procesando tareas de prioridad 2: ")
	for _, t := range tareas2 {
		wg.Add(1)
		canal2 <- t
	}
	wg.Wait()

	//prioridad 3
	fmt.Println("Procesando tareas de prioridad 3: ")
	for _, t := range tareas3 {
		wg.Add(1)
		canal3 <- t
	}
	wg.Wait()

	// Cerrar los canales
	close(canal0)
	close(canal1)
	close(canal2)
	close(canal3)
}
