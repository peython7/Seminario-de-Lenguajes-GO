package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Tarea struct {
	Prioridad int
	Valor     int
}

var (
	wg             sync.WaitGroup
	mutexArchivo0  sync.Mutex
	mutexArchivo1  sync.Mutex
	mutexAcumulado sync.Mutex
	acumulado      int
)

func sumarDigitos(n int) int { //usada en datos de prioridad 0
	suma := 0
	for n > 0 {
		suma += n % 10
		n /= 10
	}
	return suma
}

func invertirDigitos(n int) int { //usada en datos de prioridad 1
	invertido := 0
	for n > 0 {
		invertido = invertido*10 + n%10
		n /= 10
	}
	return invertido
}

func worker(id int, tareas <-chan Tarea, grupos *[4]sync.WaitGroup, file0, file1 *os.File) {
	for tarea := range tareas { // recibe tareas del canal
		fmt.Printf("Worker %d procesando tarea = Prioridad: %d, Valor: %d\n", id, tarea.Prioridad, tarea.Valor)
		switch tarea.Prioridad {
		case 0:
			resultado := sumarDigitos(tarea.Valor)
			mutexArchivo0.Lock() // bloquea el acceso a el archivo para evitar conflictos de escritura
			file0.WriteString(fmt.Sprintf("(0, %d)\n", resultado))
			mutexArchivo0.Unlock() // desbloquea el acceso a el archivo
			time.Sleep(100 * time.Millisecond)

		case 1:
			resultado := invertirDigitos(tarea.Valor)
			mutexArchivo1.Lock() // bloquea el acceso a el archivo
			file1.WriteString(fmt.Sprintf("(1, %d)\n", resultado))
			mutexArchivo1.Unlock() // desbloquea el acceso a el archivo
			time.Sleep(100 * time.Millisecond)

		case 2:
			resultado := tarea.Valor * 10
			fmt.Printf("Prioridad 2= %d multiplicado es = %d\n", tarea.Valor, resultado)
			time.Sleep(100 * time.Millisecond)

		case 3:
			mutexAcumulado.Lock() // bloquea el acceso a la variable acumulado
			acumulado += tarea.Valor
			fmt.Printf("Prioridad 3 acumulado (+%d): %d\n", tarea.Valor, acumulado)
			mutexAcumulado.Unlock() // desbloquea acceso a la variable acumulado
			time.Sleep(100 * time.Millisecond)
		}
		grupos[tarea.Prioridad].Done() //  decremento el contador de tareas pendientes
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	cantTareas := rand.Intn(81) + 20
	var tareas []Tarea

	//Creo los archivos
	file0, _ := os.Create("prioridad0.txt")
	file1, _ := os.Create("prioridad1.txt")
	defer file0.Close()
	defer file1.Close()

	for i := 0; i < cantTareas; i++ {
		valor := rand.Intn(500) + 1
		prioridad := rand.Intn(4)
		tareas = append(tareas, Tarea{prioridad, valor}) //agrego la tarea al slice de tareas
	}

	tareaCh := make(chan Tarea)

	var grupos [4]sync.WaitGroup //defino un arreglo de 4 WaitGroups (uno por prioridad)

	// Inicio a los 4 workers
	for i := 1; i <= 4; i++ {
		go worker(i, tareaCh, &grupos, file0, file1)
	}

	// Scheduler
	for prioridad := 0; prioridad <= 3; prioridad++ {
		for _, tarea := range tareas {
			if tarea.Prioridad == prioridad {
				grupos[prioridad].Add(1) //incremento el contador de tareas pendientes de esa prioridad en el arreglo grupos
				tareaCh <- tarea
			}
		}
		grupos[prioridad].Wait() //detenemos el main hasta que se terminen de procesar todas las tareas de esa prioridad
	}

	close(tareaCh)
	fmt.Println("Todas las tareas fueron procesadas.")
}
