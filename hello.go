package main

import (
	"fmt"
	"reflect"
	"os"
	// "net/http"
)

func main() {
	//exibeIntroducao()
	//exibeMenu()
	nome,idade := devolveNomeEIdade()
	fmt.Println(nome,idade)
	comando := leComando()
	
	switch comando{
	case 1:
		inicarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando.")
		os.Exit(-1)
	}

}
func devolveNomeEIdade() (string, int){
	nome := "Douglas"
	idade := 24
	return nome, idade
}

func exibeIntroducao(){
	nome := "Douglas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr", nome, "sua idade é", idade)
	fmt.Println("Esse programa está na versão", versao)
	fmt.Println("O tipo da variável é", reflect.TypeOf(versao))
}
func exibeMenu(){
	fmt.Println("1 - Iniciar monitoramento do site")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}
func leComando() int{
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func inicarMonitoramento(){
	fmt.Println("Monitorando...")
	//site := "https://gustavodscruz.dev/"
	//resp, err := http.Get(site) 
		
}
