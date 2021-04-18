package model
import "../db"

type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Qtd int
}

func BuscaTodosProdutos() []Produto{

	db:=ConectaDataBase()
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
	return produtos
}