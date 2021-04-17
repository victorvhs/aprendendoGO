package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	for {
		exibeIntroducer()	
		escolha(leComando())
		}
}

func escolha(comando int){
	switch comando{
	case 1:
		
		iniciarMonitoramento()
	case 2:
		fmt.Println("Aqui estão os logs")
	case 0:
		fmt.Println("Saindo")
		os.Exit(0)

	default:
		fmt.Println("Comando Invalido. \nEssas são as opções")
		fmt.Println("1\tIniciar Monitoramento")
		fmt.Println("2\tExibir os logs")
		fmt.Println("0\tSair")
	}
}

func exibeIntroducer(){
	versao:=1.0
	fmt.Println("###################################")
	fmt.Println( "\tVersão",versao)

	fmt.Println("#1\tIniciar Monitoramento")
	fmt.Println("#2\tExibir os logs")
	fmt.Println("#0\tSair")
	fmt.Print(":>")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func iniciarMonitoramento(){
	fmt.Println("Começando a monitorar")
	site := "https://www.alura.com.br/044"
	site = "https://random-status-code.herokuapp.com/"
	resp,_:= http.Get(site)
	if resp.StatusCode ==200{
		fmt.Println("+ Site:",site," Carregado ok")
	}else{
		fmt.Println("- Site:",site,"com erro: ",resp.Status)
	}
}