package main

import "fmt"

func main() {
	ano := 1991
	for {
		if ano >= 2020 {
			break
		}
		fmt.Println(ano)
		ano++
	}
}
