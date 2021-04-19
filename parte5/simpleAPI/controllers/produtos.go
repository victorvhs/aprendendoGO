package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/vhreis/aprendendoGO/parte5/simpleAPI/model"
)
var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	todosOsProdutos := model.BuscaTodosProdutos()
	templates.ExecuteTemplate(w,"Index",todosOsProdutos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w,"New",nil)
}

func Inserir(w http.ResponseWriter, r *http.Request){
	if r.Method ==  "POST"{
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		qtd := r.FormValue("qtd")

		precoConvertido,err := strconv.ParseFloat(preco,64)
		if err != nil{
			log.Println("Preço não valido")
		}
		qtdConvertida, err := strconv.Atoi(qtd)
		if err != nil{
			log.Println("Qtd não valido")
		}
		model.NovoProduto(nome,descricao,precoConvertido,qtdConvertida)
	}
	http.Redirect(w,r,"/",301)
}