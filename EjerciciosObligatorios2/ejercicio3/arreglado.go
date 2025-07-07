package main

import (
	"errors"
	"fmt"
)

type Elemento struct {
	numero   int
	cantidad int
}

type OptimumSlice struct {
	Info []Elemento
}

func New(slice []int) OptimumSlice {
	if len(slice) == 0 {
		return OptimumSlice{} // Retorna un OptimumSlice vacío si el slice de entrada está vacío
	}

	resultado := OptimumSlice{} 
	nroAct := slice[0]          // Toma el primer número como el número actual
	cont := 1                   // Inicializa el contador en 1

	for i := 1; i < len(slice); i++ { // Recorre el slice desde el segundo elemento
		if slice[i] == nroAct { // Si el número actual es igual al número anterior
			cont++ 
		} else {
			// Agrega el número actual con su cantidad
			resultado.Info = append(resultado.Info, Elemento{numero: nroAct, cantidad: cont})
			nroAct = slice[i]
			cont = 1
		}
	}
	// Agrega el último grupo que quedó sin guardar
	resultado.Info = append(resultado.Info, Elemento{numero: nroAct, cantidad: cont})

	return resultado 
}

func IsEmpty(o OptimumSlice) bool {
	return len(o.Info) == 0 
}


func Len(o OptimumSlice) int {
	var elementos int = 0             
	for i := 0; i < len(o.Info); i++ { 
		elementos += o.Info[i].cantidad // Suma la cantidad de cada elemento al contador
	}
	return elementos // Retorna el total de elementos contando las cantidades
}


func FrontElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, errors.New("el OptimumSlice está vacío") // Retorna un error si el OptimumSlice está vacío
	}
	return o.Info[0].numero, nil 
}


func LastElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, errors.New("el OptimumSlice está vacío") // Retorna un error si el OptimumSlice está vacío
	}

	return o.Info[len(o.Info)-1].numero, nil 
}

func insert(o OptimumSlice, v int, p int) OptimumSlice {
	var s OptimumSlice
	ocurrencias := 0

	for i := 0; i < len(o.Info); i++ {
		inicioBloque := ocurrencias                       // Posición de inicio del bloque actual
		finBloque := ocurrencias + o.Info[i].cantidad - 1 // Posición de fin del bloque actual

		if p >= inicioBloque && p <= finBloque {
			// Divido el bloque actual en 2 partes
			primerParte := p - inicioBloque
			segundaParte := o.Info[i].cantidad - primerParte

			if primerParte > 0 { // Chequeamos que hayan elementos antes de insertar 'v'
				s.Info = append(s.Info, Elemento{o.Info[i].numero, primerParte})
			}
			s.Info = append(s.Info, Elemento{v, 1})
			if segundaParte > 0 { // Si hay elementos después de la posición p
				s.Info = append(s.Info, Elemento{o.Info[i].numero, segundaParte}) // Agrega el resto del bloque actual
			}
			// Agregar el resto del slice
			s.Info = append(s.Info, o.Info[i+1:]...)
			return s
		}

		s.Info = append(s.Info, o.Info[i])
		ocurrencias += o.Info[i].cantidad
	}
	return s
}

func Insert(os *OptimumSlice, elem int, pos int) (bool, error) {
	o := (*os)
	if pos < 0 || pos > Len(o) {
		return false, errors.New("No se pudo insertar")
	}

	if pos == Len(o) { // Si la posición es igual al largo del OptimumSlice, se agrega al final directamente
		if len(o.Info) > 0 && o.Info[len(o.Info)-1].numero == elem {
			o.Info[len(o.Info)-1].cantidad++
		} else {
			o.Info = append(o.Info, Elemento{elem, 1})
		}
		*os = o
		return true, nil
	}
	arregloDevalores := SliceArray(o)
	var ocurrencias int = 0
	var valorActual int
	for i := 0; i < len(o.Info); i++ {
		valorActual = o.Info[i].numero
		for j := 0; j < o.Info[i].cantidad; j++ {
			if pos == ocurrencias {
				if valorActual == elem {
					o.Info[i].cantidad++
				} else if (i-1) >= 0 && arregloDevalores[ocurrencias-1] == elem {
					o.Info[i-1].cantidad++
				} else if (i+1) < len(o.Info) && arregloDevalores[ocurrencias+1] == elem {
					o.Info[i+1].cantidad++
				} else {
					*os = insert(o, elem, pos) // Llama a la función insert para insertar el elemento en la posición indicada y lo devuelve
				}
				return true, nil
			}
			ocurrencias++
		}
	}
	return false, errors.New("No se pudo insertar") // Retorna false si no se pudo insertar el elem en la posición indicada
}

func SliceArray(o OptimumSlice) []int {
	var slice []int
	for _, r := range o.Info {
		for i := 0; i < r.cantidad; i++ {
			slice = append(slice, r.numero)
		}
	}
	return slice
}

func main() {
	slice := []int{1, 1, 1, 1, 3, 3, 3, 3, 3, 5, 7, 7, 7, 8}
	o := New(slice)
	fmt.Println("OptimumSlice:", o)

	inserted, err := Insert(&o, 5, 6)
	if err != nil {
		fmt.Println(err) // si no se puede insertar imprime el error
	} else {
		fmt.Println("Insertado:", inserted) // imprime "insertado" si el elemento se inserta correctamente
	}

	fmt.Println("OptimumSlice:", o)
	fmt.Println("Slice Array:", SliceArray(o)) // imprime el resultado final
}
