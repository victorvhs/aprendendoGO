package main

import (
	"fmt"
)

func main() {
	dados := [][]string{
		[]string{"maria", "padilha", "dançar"},
		[]string{"Omalu", "xaxara", "Curar"},
		[]string{"Nanã", "Buluku", "Rainha"},
	}
	for _, v := range dados {
		fmt.Printf("%v tem como gosta de %v\n", v[0], v[2])
	}
	for _, v := range dados {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}
