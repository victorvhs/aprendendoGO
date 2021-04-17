package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)
const yellow = "\033[33m"
const red="\033[31m"
const green  = "\033[32m"
const purple = "\033[35m "
const reset  = "\033[0m"
const deleyInSeconds = 1
const arquivoSites = "sites.txt"

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
		imprimeLogs()
	case 3:
		fmt.Println("Novo site para monitorar")
		novoSite := leComandoString()
		gravarLinha(novoSite)
	case 0:
		fmt.Println("Saindo")
		os.Exit(0)

	default:
		fmt.Println("Comando Invalido. \nEssas são as opções")
		fmt.Println("1\tIniciar Monitoramento")
		fmt.Println("2\tExibir os logs")
		fmt.Println("3\tNovo site para monitorar")
		fmt.Println("0\tSair")
	}
}

func exibeIntroducer(){
	versao:=1.0
	fmt.Println( "\tVersão",versao)

	fmt.Println("1\tIniciar Monitoramento")
	fmt.Println("2\tExibir os logs")
	fmt.Println("3\tAdicionar novo site")
	fmt.Println(purple,"\r0\tSair",reset)
	fmt.Print(":>")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
func leComandoString() string {
	var comando string
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

	if resp.StatusCode ==200 {
		fmt.Println(green,"\r+ Site:",reset,site," Carregado ok")
		registraLog(site,true)
	}else{
		fmt.Println(red,"\r- Site:",reset,site,"com erro: ",resp.Status)
		registraLog(site,false)
	}
}

func leArquivoSites()[]string{

	var sites []string
	file,err := os.Open(arquivoSites)
	if err != nil{
		fmt.Println(red,"Error:",err,reset)
	}
	leitor := bufio.NewReader(file)
	for{
		linha, err :=leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		linha = strings.Trim(linha,"\n")
		sites = append(sites,linha)
		if err == io.EOF{
			break
		}
	}
	file.Close()
	return sites
}

func gravarLinha(linha string,){
	f, err := os.OpenFile(arquivoSites, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err !=nil {
		fmt.Println(red,"Erro:",err,reset)
	}
	linha += "\n"
	if _, err := f.WriteString(linha); err != nil {
		fmt.Println(red,"Error",err,reset)
	}
	f.Close()
}

func registraLog(site string, status bool){

	file, err := os.OpenFile("log.txt",os.O_APPEND|os.O_RDWR|os.O_CREATE,0666)
	if err != nil {
		fmt.Println(red,"\rERROR:",err,reset)
	}
	if status{
		linha := time.Now().Format("02/01/2006 15:04:05")+"-" + site+"- online: "+ strconv.FormatBool(status)+"\n"
		file.WriteString(linha)}
	linha :=time.Now().Format("02/01/2006 15:04:05")+"-"+site+"- offline: "+ strconv.FormatBool(status)+"\n"
	file.WriteString(linha)
	
	file.Close()

}

func imprimeLogs(){

	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(red,"\nError: ",err,reset)
	}
	fmt.Println(string(file))
}