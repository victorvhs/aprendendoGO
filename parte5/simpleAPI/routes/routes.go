package routes

import (
	"net/http"

	"github.com/vhreis/aprendendoGO/parte5/simpleAPI/controllers"
)

func CarregaRotas (){
	http.HandleFunc("/",controllers.Index)
	http.HandleFunc("/new",controllers.NovoProduto)
	http.HandleFunc("/insert",controllers.Inserir)
	http.HandleFunc("/delete",controllers.Deletar)
	http.HandleFunc("/edit",controllers.Editar)
	http.HandleFunc("/update",controllers.Atualizar)
}
