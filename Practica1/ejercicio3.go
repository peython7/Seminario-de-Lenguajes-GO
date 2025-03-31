package main

import "fmt"

func main() {
	/* integers */
	var zz int = 0xA
	var x int = 10
	var z int = x
	var y int8 = int8(x + 1)
	const n int = 5001
	const c int = 5001
	/* float */
	var e float32 = 6
	var f float32 = e

	fmt.Println(zz)
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(n)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)
}
