package ejercicio1

import (
	"fmt"
	"strconv"
)


type List *element

type element struct {
	val  Ingresante
	next List
}

// ----------- OPERACIONES DE LISTA ------------

func New() List {
	return nil
}

func IsEmpty(l List) bool {
	return l == nil
}

func FrontElement(l List) Ingresante {
	if l == nil {
		panic("Lista vacía")
	}
	return l.val
}

func Next(l List) List {
	if l == nil {
		panic("Lista vacía")
	}
	return l.next
}

func PushBack(l *List, elem Ingresante) {
	if *l == nil {
		*l = &element{val: elem}
		return
	}
	actual := *l
	for actual.next != nil {
		actual = actual.next
	}
	actual.next = &element{val: elem}
}

// ----------- LÓGICA DE LA CONSIGNA ------------

func ProcesarIngresantes(l *List) {
	if IsEmpty(*l) {
		fmt.Println("Lista vacía.")
		return
	}

	nacimientosPorAnio := make(map[int]int)
	inscriptosPorCarrera := make(map[string]int)

	actual := *l
	var anterior *element = nil

	for actual != nil {
		i := actual.val

		// a) Ciudad: Bariloche
		if i.Ciudad == "Bariloche" {
			fmt.Println("De Bariloche:", i.Nombre, i.Apellido)
		}

		// b) Año de nacimiento
		nacimientosPorAnio[i.Nacimiento.Anio]++

		// c) Carrera
		inscriptosPorCarrera[i.Carrera]++

		// d) Eliminar si no presentó título
		if !i.TituloPresentado {
			if anterior == nil {
				*l = actual.next
				actual = *l
				continue
			} else {
				anterior.next = actual.next
				actual = anterior.next
				continue
			}
		}

		anterior = actual
		actual = actual.next
	}

	// b) Año con más nacimientos
	var anioMax int
	var maxNac int
	for anio, cant := range nacimientosPorAnio {
		if cant > maxNac {
			anioMax = anio
			maxNac = cant
		}
	}
	fmt.Println("Año con más nacimientos:", anioMax)

	// c) Carrera con más inscriptos
	var carreraMax string
	var maxCarrera int
	for carrera, cant := range inscriptosPorCarrera {
		if cant > maxCarrera {
			carreraMax = carrera
			maxCarrera = cant
		}
	}
	fmt.Println("Carrera con más inscriptos:", carreraMax)
}

// ----------- CONVERSIÓN A STRING PARA IMPRIMIR ------------

func ToString(i Ingresante) string {
	return i.Nombre + " " + i.Apellido + " - " + i.Ciudad + " - " +
		strconv.Itoa(i.Nacimiento.Dia) + "/" +
		strconv.Itoa(i.Nacimiento.Mes) + "/" +
		strconv.Itoa(i.Nacimiento.Anio) + " - " +
		"Carrera: " + i.Carrera
}

// ----------- FUNCIÓN PRINCIPAL ------------

func main() {
	l := New()

	PushBack(&l, Ingresante{"Gómez", "Ana", "Bariloche", FechaNacimiento{10, 5, 2002}, true, "APU"})
	PushBack(&l, Ingresante{"Pérez", "Luis", "Neuquén", FechaNacimiento{3, 7, 2001}, false, "LI"})
	PushBack(&l, Ingresante{"López", "Carla", "Bariloche", FechaNacimiento{20, 10, 2002}, true, "LS"})
	PushBack(&l, Ingresante{"Martínez", "Sofía", "Cipolletti", FechaNacimiento{15, 8, 2001}, true, "APU"})
	PushBack(&l, Ingresante{"Ramos", "Julián", "Roca", FechaNacimiento{1, 1, 2002}, false, "LS"})

	fmt.Println("Procesando ingresantes...\n")
	ProcesarIngresantes(&l)

	fmt.Println("\nLista final (sólo con quienes presentaron título):")
	actual := l
	for actual != nil {
		fmt.Println(ToString(actual.val))
		actual = actual.next
	}
}
