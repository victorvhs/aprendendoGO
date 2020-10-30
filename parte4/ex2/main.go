package main

import (
	"fmt"
)

func main() {
	meses := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range meses {
		fmt.Println(v)
	}
	fmt.Printf("O tipo do Slice é: %T", meses)
}
