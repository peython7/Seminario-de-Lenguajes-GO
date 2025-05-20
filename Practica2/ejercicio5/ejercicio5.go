package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	elementos = 5
)

type Vector [elementos]float64

func Initialize(v *Vector, f float64) {
	for i := 0; i < elementos; i++ {
		v[i] = math.Round(rand.Float64() * 10)
	}
}

func Sum(v1, v2 Vector) Vector {
	var v3 Vector
	for i := 0; i < elementos; i++ {
		v3[i] = v1[i] + v2[i]
	}
	return v3
}

func SumInPlace(v1 *Vector, v2 Vector) {
	for i := 0; i < elementos; i++ {
		v1[i] += v2[i]
	}
}

func main() {
	var vector1 Vector
	var vector2 Vector
	var f float64
	Initialize(&vector1, f)
	fmt.Println(vector1)
	Initialize(&vector2, f)
	fmt.Println(vector2)
	vector3 := Sum(vector1, vector2)
	fmt.Println(vector3)

	SumInPlace(&vector1, vector2)
	fmt.Println(vector1)
}
