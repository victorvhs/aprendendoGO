package main

import (
	"fmt"
)

func main() {
	dados := map[string][]string{
		"Lucas": []string{
			"ler", "dormir", "comer",
		},
		"Lurdes": []string{
			"cozinhar", "estudar",
		},
		"Leticia": []string{
			"correr", "bicicleta", "nadar",
		},
	}
	dados["Luiz"] = []string{"Jogar", "Cinema", "Tocar banjo"}

	for k, v := range dados {
		fmt.Printf("%v gosta de: \n", k)
		for i, t := range v {
			fmt.Println("\t", i, "#", t)
		}
	}
}
