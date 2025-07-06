package ejercicio1 
import "fmt"

type Nodo struct {
	data      Estudiante
	siguiente *Nodo
}

type List struct {
	head *Nodo
}

func New() List {
	return List{head: nil}
}

// Retorna true si la lista está vacía.
func IsEmpty(l *List) bool {
	return l.head == nil
}

// agrega un nuevo nodo al final de la lista.
func (l *List) Agregar(valor Estudiante) {
	nuevoNodo := &Nodo{data: valor}
	if IsEmpty(l) {
		l.head = nuevoNodo
	} else {
		actual := l.head
		for actual.siguiente != nil {
			actual = actual.siguiente
		}
		actual.siguiente = nuevoNodo
	}
}

func FrontElement(l List) Estudiante {
	if IsEmpty(l){
		fmt.Println("lista vacía") // error si la lista no tiene nodos
	}
	return l.head.data
}

func Len(l List) int {
	if IsEmpty(l) {
		return 0
	} else {
		cant := 0
		actual := l.head
		for actual != nil {
			cant++
			actual = actual.siguiente
		}
		return cant
	}
}

func ToString(l List) string {
	if IsEmpty(l) {
		return "La lista está vacía"
	}
	actual := l.head
	var resultado string
	for actual != nil {
		resultado += fmt.Sprintf("%v ", actual.data)
		actual = actual.siguiente
	}
	return resultado

}

func PushFront(l *List, valor Estudiante) {
	nuevoNodo := &Nodo{data: valor}
	if IsEmpty(l) {
		l.head = nuevoNodo
	} else {
		nuevoNodo.siguiente = l.head
		l.head = nuevoNodo
	}
}

func PushBack(l *List, valor Estudiante) {
	nuevoNodo := &Nodo{data: valor}
	if IsEmpty(l) {
		l.head = nuevoNodo
	} else {
		actual := l.head
		for actual.siguiente != nil {
			actual = actual.siguiente
		}
		actual.siguiente = nuevoNodo
	}
}

func (l List) Iterate() {
	if IsEmpty(&l) {
		fmt.Println("La lista está vacía")
		return
	}
	actual := l.head
	for actual != nil {
		fmt.Println(actual.data)
		actual = actual.siguiente
	}
}
