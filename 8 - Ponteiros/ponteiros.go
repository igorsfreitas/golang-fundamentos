package main

import "fmt"

func main() {
	// PONTEIROS SÃO UMA VARIÁVEL QUE GUARDAM O ENDEREÇO DE MEMÓRIA DE OUTRA VARIÁVEL, E NÃO O SEU VALOR

	fmt.Println("PONTEIROS")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2) // NESSE CASO A VARIÁVEL 2 NÃO É ALTERADA POIS ELA RECEBEU O VALOR DA VARIÁVEL 1, E NÃO O ENDEREÇO DE MEMÓRIA

	// PONTEIROS SÃO REPRESENTADOS POR UM *
	// PONTEIROS SÃO UMA REFERENCIA DE MEMÓRIA
	var variavel3 int = 100
	var ponteiro *int

	ponteiro = &variavel3 // COM O & É POSSÍVEL PEGAR O ENDEREÇO DE MEMÓRIA DE UMA VARIÁVEL

	fmt.Println(variavel3, ponteiro) // NESSE CASO O VALOR DA VARIAVEL PONTEIRO É O ENDEREÇO DE MEMÓRIA DA VARIÁVEL 3 EX: 0xc0000b4008
	fmt.Println(*ponteiro)           // DESREFERENCIAÇÃO - NESSE CASO O VALOR DA VARIAVEL PONTEIRO É O VALOR DA VARIÁVEL 3 EX: 100
}
