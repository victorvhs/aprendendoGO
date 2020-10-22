package main

import (
	"fmt"
)

func main() {
	x := 0
	if x > 10 {
		fmt.Println("Nota 100")
	} else if x <= 0 {
		fmt.Println("Sem nota")
	}
}
