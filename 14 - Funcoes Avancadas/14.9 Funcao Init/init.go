package main

import "fmt"

// QUANDO PRECISA SETAR ALGUM VALOR,
// OU FAZER ALGUMA CONFIGURAÇÃO INICIAL
// USA-SE A FUNCAO INIT
var n int

func init() {
	fmt.Println("Função Init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
