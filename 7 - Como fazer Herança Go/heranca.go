package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa // AQUI É BASICAMENTE COMO SE EU TIVESSE COPIANDO O CAMPO DA STRUCT PESSOA PARA A STRUCT ESTUDANTE, CASO EU QUEIRA CRIAR UM
	// CAMPO PESSOA DENTRO DA STRUCT ESTUDANTE, EU PRECISO DECLARAR O TIPO DA STRUCT PESSOA DENTRO DA STRUCT ESTUDANTE
	// EXEMPLO: pessoa pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"} // AQUI EU ESTOU PASSANDO A STRUCT PESSOA COMO UM CAMPO DA STRUCT ESTUDANTE
	fmt.Println(e1.nome)                     // CASO EU QUEIRA ACESSAR UM CAMPO DA STRUCT PESSOA DENTRO DA STRUCT ESTUDANTE, EU PRECISO ACESSAR O CAMPO
}
