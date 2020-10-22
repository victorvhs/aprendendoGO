package main

import "fmt"

type repolho int

func main() {
	var x repolho
	fmt.Printf("X valor : %v\n", x)
	x = 42
	fmt.Printf("X valor : %v\nTipo de X : %T\n", x, x)
}
