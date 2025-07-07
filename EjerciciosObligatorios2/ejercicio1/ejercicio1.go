package ejercicio1

import "fmt"

func informarBariloche(lista List) {
	for actual := lista.head; actual != nil; actual = actual.siguiente {
		if actual.data.ciudad == "Bariloche" {
			fmt.Println(actual.data.nombre, "", actual.data.apellido)
		}
	}
}

func calcularAnio(lista List) int {
	contador := make(map[int]int)
	for actual := lista.head; actual != nil; actual = actual.siguiente {
		contador[actual.data.fecha.anio]++
	}
	max := 0
	anio := 0
	for k, v := range contador {
		if v > max {
			max = v
			anio = k
		}
	}
	return anio
}

func calcularCarrera(lista List) string {
	contador := make(map[int]int)
	for actual := lista.head; actual != nil; actual = actual.siguiente {
		contador[actual.data.codigo]++
	}
	max := 0
	codigo := 0
	for k, v := range contador {
		if v > max {
			max = v
			codigo = k
		}
	}
	switch codigo {
	case 1:
		return "APU"
	case 2:
		return "LI"
	case 3:
		return "LS"
	default:
		return "No se encontró la carrera"
	}
}

func EliminarSinTitulo(lista *List) {
	// Eliminar nodos desde el principio mientras el título sea falso
	for lista.head != nil && !lista.head.data.titulo {
		lista.head = lista.head.siguiente
	}

	// Si la lista quedó vacía, no hay nada más que hacer
	if lista.head == nil {
		return
	}

	// Recorrer la lista a partir del primer nodo con título
	actual := lista.head //actual es un puntero al mismo nodo que lista.head
	for actual != nil && actual.siguiente != nil {
		if !actual.siguiente.data.titulo {
			// Saltar el nodo sin título
			actual.siguiente = actual.siguiente.siguiente
		} else {
			actual = actual.siguiente
		}
	}
}

func main() {
	lista := New()
	
	estudiante1 := Estudiante{
		nombre:   "Juan",
		apellido: "Pérez",
		ciudad:   "La Plata",
		fecha: FechaNacimiento{
			dia:  15,
			mes:  3,
			anio: 2000,
		},
		titulo: true,
		codigo: 1, // APU
	}

	estudiante2 := Estudiante{
		nombre:   "María",
		apellido: "González",
		ciudad:   "Bariloche",
		fecha: FechaNacimiento{
			dia:  22,
			mes:  7,
			anio: 1998,
		},
		titulo: false,
		codigo: 1, // APU
	}

	estudiante3 := Estudiante{
		nombre:   "Carlos",
		apellido: "López",
		ciudad:   "Rosario",
		fecha: FechaNacimiento{
			dia:  10,
			mes:  11,
			anio: 2000,
		},
		titulo: true,
		codigo: 3, // LS
	}
	informarBariloche(lista)
	fmt.Println("El año que mas ingresantes nacieron fue", calcularAnio(lista))
	fmt.Println("La carrera con mas ingresantes inscriptos fue", calcularCarrera(lista))
	lista.Iterate()
	EliminarSinTitulo(&lista)
	fmt.Println("Lista nueva solo con los ingresantes que tienen titulo")
	fmt.Println(ToString(lista))
}
