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
	return o[len(o)-1].Valor, nil //Si devuelve nil es porque no hubo error
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
	var posicionFinal int = 0
	var posicionInicial int = 0
	var ocurrenciasAnterior int = 0
	for i := 0; i < len(o); i++ {
		posicionInicial = ocurrenciasAnterior
		posicionFinal = posicionInicial + (o[i].Ocurrencias - 1)
		if p >= posicionInicial && p <= posicionFinal {
			if p == posicionInicial {
				s = append(s, Racha{v, 1})
				s = append(s, o[i])
			} else if p == posicionFinal {
				s = append(s, o[i])
				s = append(s, Racha{v, 1})
			} else {
				OcurrenciasRacha1 := p - posicionInicial
				OcurrenciasRacha2 := (posicionFinal + 1) - p
				s = append(s, Racha{o[i].Valor, OcurrenciasRacha1})
				s = append(s, Racha{v, 1})
				s = append(s, Racha{o[i].Valor, OcurrenciasRacha2})
			}
		} else {
			s = append(s, o[i])
		}
		ocurrenciasAnterior += o[i].Ocurrencias
	}
	return s
}

func Insert(o *OptimumSlice, v int, p int) int {
	racha := (*o)
	if p < 0 && p >= Len(racha) {
		return -1
	}
	var ocurrencias int = 0
	var valorActual int
	for i := 0; i < len(racha); i++ {
		valorActual = racha[i].Valor
		for j := 0; j < racha[i].Ocurrencias; j++ {
			if p == ocurrencias {
				if valorActual == v {
					racha[i].Ocurrencias++
				} else if (i-1) >= 0 && racha[i-1].Valor == v {
					racha[i-1].Ocurrencias++
				} else if (i+1) < len(racha) && racha[i+1].Valor == v {
					racha[i+1].Ocurrencias++
				} else {
					*o = insert(racha, v, p)
				}
				return 0
			}
			ocurrencias++
		}
	}
	return -1
}

func main() {
	s := []int{1, 1, 1, 3, 3, 3, 3, 3, 5, 7, 7, 7}
	os := New(s)
	fmt.Println("Resultado de ejecutar el New: \n", os)

	Insert(&os, 8, 11)
	fmt.Println("Resultado de ejecutar el Insert: \n", os)

	var err error
	var primerElemento int
	primerElemento, err = FrontElement(os)
	if err == nil {
		fmt.Println("El primer elemento es: ", primerElemento)
	} else {
		fmt.Println(err)
	}

	var ultimoElemento int
	ultimoElemento, err = LastElement(os)
	if err == nil {
		fmt.Println("El ultimo elemento es: ", ultimoElemento)
	} else {
		fmt.Println(err)
	}

	fmt.Println("La cantidad de elementos en el arreglo ", os, " es: ", Len(os))

	var array []int = SliceArray(os)
	fmt.Println("El arreglo de ", os, " es: ", array)

}
