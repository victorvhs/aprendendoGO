package main

import (
	"fmt"
)

func main() {
	divida := [5]int{150, 120, 200, 30, 50}
	for indice, valor := range divida {
		fmt.Printf("Indice: %v = %v\n", indice, valor)
	}
	fmt.Printf("O tipo do array é %T\n", divida)
}
