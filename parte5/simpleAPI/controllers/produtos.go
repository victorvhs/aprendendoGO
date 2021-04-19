package controllers

import (
	"net/http"
	"text/template"

	"github.com/vhreis/aprendendoGO/parte5/simpleAPI/model"
)
var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	todosOsProdutos := model.BuscaTodosProdutos()
	templates.ExecuteTemplate(w,"Index",todosOsProdutos)
}