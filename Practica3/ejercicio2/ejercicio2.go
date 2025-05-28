package main

import (
	"errors"
	"fmt"
	"strings"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func New[T any]() List[T] {
	return List[T]{}
}

func IsEmpty[T any](l List[T]) bool {
	return l.head == nil
}

func Len[T any](l List[T]) int {
	cont := 0
	for i := l.head; i != nil; i = i.next {
		cont++
	}
	return cont
}

func FrontElement[T any](l List[T]) (T, error) {
	var aux T
	if IsEmpty(l) {
		return aux, errors.New("La lista esta vacia")
	}
	return l.head.val, nil
}

func Next[T any](l List[T]) List[T] {
	if IsEmpty(l) {
		return List[T]{}
	}
	return List[T]{head: l.head.next}
}

// Agregar al principio de la lista
func PushFront[T any](l *List[T], v T) {
	e := &element[T]{val: v, next: l.head}
	l.head = e
	if l.tail == nil {
		l.tail = e
	}
}

// Agregar al final de la lista
func PushBack[T any](l *List[T], v T) {
	e := &element[T]{val: v}
	if l.tail == nil {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		l.tail = e
	}
}

// Eliminar el primer elemento y devolverlo
func Remove[T any](l *List[T]) (T, error) {
	var aux T
	if IsEmpty(*(l)) {
		return aux, errors.New("La lista esta vacia")
	}
	valor := l.head.val
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return valor, nil
}

func ToString[T any](l List[T]) string {
	if IsEmpty(l) {
		return "Lista vacia"
	} else {
		var aux strings.Builder
		aux.WriteString("{")
		for e := l.head; e != nil; e = e.next {
			fmt.Fprintf(&aux, "%v", e.val)
			if e.next != nil {
				aux.WriteString(", ")
			}
		}
		aux.WriteString("}")
		return aux.String()
	}
}

//Preguntar por el Iterate
// (que deberia hacer y que deberia devolver en caso de que devuelva algo)

func main() {
	listaEnteros := New[int]()
	PushBack(&listaEnteros, 7)
	PushBack(&listaEnteros, 17)
	PushFront(&listaEnteros, 45)
	PushBack(&listaEnteros, 9)
	PushFront(&listaEnteros, 23)
	fmt.Println(ToString(listaEnteros))
	Remove(&listaEnteros)
	fmt.Println("Aplicamos el Remove y: ")
	fmt.Println(ToString(listaEnteros))
	primerElemento, err := FrontElement(listaEnteros)
	if err == nil {
		fmt.Println("El primer elemento es: ", primerElemento)
	} else {
		fmt.Println(err)
	}

}
