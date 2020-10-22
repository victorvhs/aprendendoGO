package main

import (
	"fmt"
)

func main() {
	x := 10
	switch {
	case x == 10:
		fmt.Println("10")
	case x < 10:
		fmt.Print("menor")
	}
}
