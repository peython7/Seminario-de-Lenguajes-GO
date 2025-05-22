package main

type Racha struct {
	Valor       int
	Ocurrencies int
}

type OptimumSlice []Racha

func insert(o OptimumSlice, v int, p int)OptimumSlice{
	var s OptimumSlice
	var posicionFinal int = 0
	var posicionInicial int = 0
	var ocurrenciasAnterior int = 0
	for i:=0;i<len(o);i++{
		posicionInicial= ocurrenciasAnterior
		posicionFinal = posicionInicial + (o[i].Ocurrencies -1)
		if p >= posicionInicial && p <=posicionFinal{
			if p == posicionInicial{
				s = append(s, Racha{v,1})
				s = append(s, o[i])
			}else if p == posicionFinal{
				s = append(s, o[i])
				s = append(s, Racha{v,1})
			}else{
				OcurrenciasRacha1:=  p - posicionInicial
				OcurrenciasRacha2:= (posicionFinal + 1) - p
				s = append(s, Racha{o[i].Valor,OcurrenciasRacha1})
				s = append(s, Racha{v,1})
				s = append(s, Racha{o[i].Valor,OcurrenciasRacha2})
			}
		}else{
			s = append(s, o[i])
		}
		ocurrenciasAnterior+=o[i].Ocurrencias
	}
}
racha 1 = p - posicionInicial
racha 2 = posicionFinal -p

{ Valor: 3, Ocurrencies: 5 },
  { Valor: 2, Ocurrencies: 3 },
  { Valor: 3, Ocurrencies: 6 },


func Insert(array OptimumSlice, valor int, posicion int) int {
	if posicion < 0 && posicion >= Len(array) {
		return -1
	}
	var racha Racha
	var posicionInicial int = 0
	var posicionFinal int = 0
	var ocurrenciasAnterior int = 0
	for i := 0; i < len(array); i++ {
		racha = array[i]
		posicionInicial = ocurrenciasAnterior
		posicionFinal = posicionInicial + (racha.Ocurrencies - 1)
		if posicion >= posicionInicial && posicion <= posicionFinal {
			// La porcion en la que estoy parado es valida
			if valor == racha.Valor {
				racha.Ocurrencies++
			} else if i != 0 && array[i-1].Valor == valor {
				array[i-1].Ocurrencies++
			} else if i < len(array)-1 && array[i+1].Valor == valor {
				array[i+1].Ocurrencies++
			} else {
				var array2 []Racha = insert(array, valor, posicion)
			}
			return 0
		}
		ocurrenciasAnterior += racha.Ocurrencies
	}
}



  posicionInicial = 6
  posicionFinal = 11
  ocurrenciasAnterior = 12

func main() {
	var s OptimumSlice
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
