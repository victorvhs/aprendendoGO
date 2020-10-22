package main

import "fmt"

func main() {
	var a int = 1
	fmt.Printf("%v,%b,%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%v,%b,%#x\n", b, b, b)

}
