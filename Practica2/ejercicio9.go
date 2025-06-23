package main

import (
	"fmt"
	"strconv"
)

// Definición de nodo y lista
type List *element

type element struct {
	val  int
	next List
}

// Crear una nueva lista vacía
func New() List {
	return nil
}

// Verifica si la lista está vacía
func IsEmpty(l List) bool {
	return l == nil
}

// Calcula la longitud de la lista
func Len(l List) int {
	count := 0
	for l != nil {
		count++
		l = l.next
	}
	return count
}

// Devuelve el primer elemento (head)
func FrontElement(l List) int {
	if l == nil {
		panic("Lista vacía")
	}
	return l.val
}

// Devuelve la sublista sin el primero
func Next(l List) List {
	if l == nil {
		panic("Lista vacía")
	}
	return l.next
}

// Agrega un elemento al frente
func PushFront(l *List, elem int) {
	newElem := &element{val: elem, next: *l}
	*l = newElem
}

// Agrega un elemento al final
func PushBack(l *List, elem int) {
	if *l == nil {
		PushFront(l, elem)
		return
	}
	actual := *l
	for actual.next != nil {
		actual = actual.next
	}
	actual.next = &element{val: elem}
}

// Remueve el primer elemento y lo devuelve
func Remove(l *List) int {
	if *l == nil {
		panic("Lista vacía")
	}
	valor := (*l).val
	*l = (*l).next
	return valor
}

// Aplica una función a cada elemento de la lista (in place)
func Iterate(l List, f func(int) int) {
	for l != nil {
		l.val = f(l.val)
		l = l.next
	}
}

// Convierte la lista en string (para imprimir)
func ToString(l List) string {
	s := "["
	for l != nil {
		s += strconv.Itoa(l.val)
		if l.next != nil {
			s += ", "
		}
		l = l.next
	}
	s += "]"
	return s
}
func main() {
	l := New()

	fmt.Println("Lista vacía:", IsEmpty(l))
	PushBack(&l, 10)
	PushBack(&l, 20)
	PushFront(&l, 5)
	PushBack(&l, 30)
	fmt.Println("Lista actual:", ToString(l))
	fmt.Println("Longitud:", Len(l))
	fmt.Println("Primer elemento:", FrontElement(l))

	Iterate(l, func(x int) int {
		return x * 2
	})
	fmt.Println("Lista después de multiplicar x2:", ToString(l))

	fmt.Println("Removido:", Remove(&l))
	fmt.Println("Lista final:", ToString(l))
}
