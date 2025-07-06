package ejercicio1

type FechaNacimiento struct {
	dia int
 	mes int
  	anio int
}

type Estudiante struct {
	nombre   string
	apellido string
	ciudad   string
	fecha    FechaNacimiento
	titulo   bool
	codigo   int //1 APU 2 LI 3 LS
}
