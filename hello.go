package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Douglas"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr", nome, "sua idade é", idade)
	fmt.Println("Esse programa está na versão", versao)
	fmt.Println("O tipo da variável é", reflect.TypeOf(versao))

	fmt.Println("1 - Iniciar monitoramento do site")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1{
		fmt.Println("Monitorando...")
	} else if comando == 2{
		fmt.Println("Exibindo logs...")
	} else if comando == 0{
		fmt.Println("Saindo do programa...")
	}else{
		fmt.Println("Não conheço este comando.")
	}

}
