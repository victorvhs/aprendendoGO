package main

import "fmt"

var esporteFavorito string

func main() {
	switch esporteFavorito {
	case "volei":
		fmt.Println("VOlei")
	case "baseaball":
		fmt.Println("baseball")
	default:
		fmt.Println("Vc gosta de coisa diferentes")
	}
}
