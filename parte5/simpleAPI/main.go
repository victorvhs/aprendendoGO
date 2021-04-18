package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))
type Produto struct{
	Nome string
	Descricao string
	Preco float64
	Qtd int
}

func main(){
	http.HandleFunc("/",index)
	http.ListenAndServe(":8000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	produtos := []Produto{
		{Nome:"carro",
		Descricao: "Um carro preto",
		Preco:150000,
		Qtd: 5},
		
		{"Apartemento",
		"Pequeno",
		980000,
		1},
		{"Fone","barulhento",15,5},

	}
	
	templates.ExecuteTemplate(w,"Index",produtos)

}