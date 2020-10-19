package main

import "fmt"

type repolho int

var y int

func main() {
	var x repolho
	fmt.Printf("X: %v\n", x)
	x = 42
	fmt.Printf("X: %v\nX Tipo %T\n", x, x)

	y = int(x)
	fmt.Printf("Y:\t%v\nY tipo:\t%T\n", y, y)

}
