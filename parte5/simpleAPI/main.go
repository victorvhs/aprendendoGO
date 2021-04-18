package main

import (
	"fmt"
	"net/http"
	"text/template"
	"./model"

)

var templates = template.Must(template.ParseGlob("templates/*.html"))


func main(){
	fmt.Println("Online: http://localhost:8000")
	http.HandleFunc("/",index)
	http.ListenAndServe(":8000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	todosOsProdutos := BuscaTodosProdutos()
	templates.ExecuteTemplate(w,"Index",todosOsProdutos)

}