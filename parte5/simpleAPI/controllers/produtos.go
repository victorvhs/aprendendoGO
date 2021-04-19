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

func Deletar(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	model.DeletaProduto(id)
	http.Redirect(w,r,"/",301)
}
func Editar(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	produto := model.EditarProduto(id)
	templates.ExecuteTemplate(w,"Edit",produto)
}

func Atualizar(w http.ResponseWriter, r *http.Request){
	if r.Method =="POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		qtd := r.FormValue("qtd")

		idInt, err := strconv.Atoi(id)
		if err != nil{
			log.Println("Erro no id")
		}
		
		precoFloat, err := strconv.ParseFloat(preco,64)
		if err != nil{
			log.Println("Erro no Preço")
		}

		qtdnt, err := strconv.Atoi(qtd)
		if err != nil{
			log.Println("Erro no qtd")
		}
		model.AtualizaProduto(idInt,nome,descricao,precoFloat,qtdnt)

		http.Redirect(w,r,"/",301)


	}
}