package main

import (
	"fmt"
	"net/http"

	"github.com/vhreis/aprendendoGO/parte5/simpleAPI/routes"
)

func main(){
	routes.CarregaRotas()
	fmt.Println("Online: http://localhost:8000")
	http.ListenAndServe(":8000",nil)
}

