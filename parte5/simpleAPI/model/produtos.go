package model

import (
	"log"

	"github.com/vhreis/aprendendoGO/parte5/simpleAPI/db"
)

type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Qtd int
}

func BuscaTodosProdutos() []Produto{

	db:=db.ConectaDataBase()
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Qtd= qtd
		produtos = append(produtos, p)
	}
	
	defer db.Close()
	return produtos
}

func NovoProduto(nome,descricao string ,preco float64,qtd int){
	db := db.ConectaDataBase()
	query, err := db.Prepare("insert into produtos(nome, descricao, preco, qtd) values($1,$2,$3,$4)")
	if err != nil {
		log.Println(err.Error())
	}
	query.Exec(nome,descricao,preco,qtd)
	defer db.Close()
}
func DeletaProduto(id string) {
	db := db.ConectaDataBase()

	query, err := db.Prepare("delete from produtos where id = $1")
	if err != nil{
		log.Println(err.Error())
	}
	query.Exec(id)
	db.Close()
}