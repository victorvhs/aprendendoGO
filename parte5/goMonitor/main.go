package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)
const yellow = "\033[33m"
const red="\033[31m"
const green  = "\033[32m"
const purple = "\033[35m "
const reset  = "\033[0m"
const deleyInSeconds = 1

func main(){
	for {
		exibeIntroducer()	
		escolha(leComando())
		}
}

func escolha(comando int){
	switch comando{
	case 1:
		fmt.Println("Quantas vezes iremos testar?>")
		qtdTest:= leComando()
		iniciarMonitoramento(qtdTest)
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
	fmt.Println( "\tVersão",versao)

	fmt.Println("1\tIniciar Monitoramento")
	fmt.Println("2\tExibir os logs")
	fmt.Println(purple,"\r0\tSair",reset)
	fmt.Print(":>")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento(qdtTest int){
	fmt.Println("Começando a monitorar")
	// sites := []string{"https://www.alura.com.br/",
	// 									"https://random-status-code.herokuapp.com/",
	// 									"https://herokuapp.com/",
	// 							}

	sites := 	leArquivoSites()	

  for i:=1; i<=qdtTest; i++{
		fmt.Println(yellow,"\rCICLO: ",i,reset)
		for  _,site := range sites{
			testSite(site)
		}
		time.Sleep(deleyInSeconds*time.Second)
		fmt.Println()
	}
	fmt.Println()
}

func testSite(site string){
	resp, err:= http.Get(site)
	if err != nil{
		fmt.Println(red,"Error:",err,reset)
	}

	if resp.StatusCode ==200{
		fmt.Println(green,"\r+ Site:",reset,site," Carregado ok")
	}else{
		fmt.Println(red,"\r- Site:",reset,site,"com erro: ",resp.Status)
	}
}

func leArquivoSites()[]string{

	var sites []string
	file,err := os.Open("sites.txt")
	if err != nil{
		fmt.Println(red,"Error:",err,reset)
	}
	leitor := bufio.NewReader(file)
	for{
		linha, err :=leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites,linha)
		if err == io.EOF{
			break
		}
	}
	file.Close()
	return sites
}