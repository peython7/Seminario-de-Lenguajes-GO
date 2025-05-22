package main

import (
	"errors"
	"fmt"
)

type Racha struct {
	Valor       int
	Ocurrencias int
}

type OptimumSlice []Racha

// Crea un nuevo OptimumSlice a partir de un slice de enteros.
func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return nil
	}
	var resultado OptimumSlice
	var actual int
	var contador int
	var i int = 0
	for i < len(s) {
		actual = s[i]
		contador = 0
		for i < len(s) && actual == s[i] {
			contador++
			i++
		}
		resultado = append(resultado, Racha{actual, contador})
	}
	return resultado
}

// Retorna true si el OptimumSlice está vacío.
func IsEmpty(o OptimumSlice) bool {
	return Len(o) == 0
}

// retorna la cantidad de elementos
// teniendo en cuenta las repeticiones
func Len(o OptimumSlice) int {
	var elementos int = 0
	for i := 0; i < len(o); i++ {
		elementos += o[i].Ocurrencias
	}
	return elementos
}

func FrontElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, errors.New("El slice esta vacio")
	}
	return o[0].Valor, nil //Si devuelve nil es porque no hubo error
}

func LastElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, errors.New("El slice esta vacio")
	}
	return o[Len(o)-1].Valor, nil //Si devuelve nil es porque no hubo error
}

// Convierte el OptimumSlice en un slice de enteros
func SliceArray(o OptimumSlice) []int {
	var sliceInt []int
	for _, r := range o {
		for i := 0; i < r.Ocurrencias; i++ {
			sliceInt = append(sliceInt, r.Valor)
		}
	}
	return sliceInt
}

func insert(o OptimumSlice, v int, p int) OptimumSlice {
	var s OptimumSlice
	var valor int
	var ocurrencias int = 0
	var contador int = 0
	var i int = 0
	for i < len(o) {
		if p == 0 {
			s = append(s, Racha{v, 1})
		}
		if ocurrencias < p {
			valor = o[i].Valor
			var j int = 0
			for ocurrencias < p && j < o[i].Ocurrencias {
				ocurrencias++
				contador++
				j++
			}
			s = append(s, Racha{valor, contador})
			if ocurrencias == p {
				s = append(s, Racha{v, 1})
			}
		} else {
			s = append(s, o[i])
		}
		i++
	}
	return s
}

func Insert(o OptimumSlice, v int, p int) int {
	if p < 0 && p >= Len(o) {
		return -1
	}
	var ocurrencias int = 0

	var valorActual int
	for i := 0; i < len(o); i++ {
		valorActual = o[i].Valor
		for j := 0; j < o[i].Ocurrencias; j++ {
			if p == ocurrencias {
				if valorActual == v {
					o[i].Ocurrencias++
				} else if (i-1) >= 0 && o[i-1].Valor == v {
					o[i-1].Ocurrencias++
				} else if o[i+1].Valor == v {
					o[i+1].Ocurrencias++
				} else {
					o = insert(o, v, p)
				}
			}
			ocurrencias++
		}
	}
	return 0
}

func main() {
	s := []int{1, 1, 1, 3, 3, 3, 3, 3, 5, 7, 7, 7}
	os := New(s)
	fmt.Println(os)
	fmt.Println(Len(os))
}
