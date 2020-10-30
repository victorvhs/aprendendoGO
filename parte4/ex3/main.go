package main

import (
	"fmt"
)

func main() {
	meses := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(meses[0:3])
	fmt.Println(meses[4:])
	fmt.Println(meses[1:7])
	fmt.Println(meses[2:9])
	fmt.Println(meses[2 : len(meses)-1])

}
