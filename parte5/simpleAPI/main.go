package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)
func conectaDataBase() * sql.DB {
	con := "user=postgres dbname=simpleAPI password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres",con)
	if err !=nil{
		fmt.Println(err, "Error")
		panic(err.Error())
		
	}
	return db
}

var templates = template.Must(template.ParseGlob("templates/*.html"))
type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Qtd int
}

func main(){
	fmt.Println("Online: http://localhost:8000")
	http.HandleFunc("/",index)
	http.ListenAndServe(":8000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	db:=conectaDataBase()
	selectProdutos,err := db.Query("select * from produtos")
	if err != nil{
		panic(err.Error())
	
	}
	p :=Produto{}
	produtos := []Produto{}
	
	for selectProdutos.Next(){
		var id,qtd int
		var nome, descricao string
		var preco float64
		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &qtd)
		if err != nil{
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Qtd= qtd
		produtos = append(produtos, p)
	}
	
	defer db.Close()
	
	templates.ExecuteTemplate(w,"Index",produtos)

}