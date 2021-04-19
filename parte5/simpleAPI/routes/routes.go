package routes

import (
	"net/http"

	"github.com/vhreis/aprendendoGO/parte5/simpleAPI/controllers"
)

func CarregaRotas (){
	http.HandleFunc("/",controllers.Index)
}
