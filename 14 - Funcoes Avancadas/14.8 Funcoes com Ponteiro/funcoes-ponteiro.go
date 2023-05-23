package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

// AQUI TO RECEBENDO UM PONTEIRO DE PARAMETRO
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	// QUANDO VC QUER FAZER ALTERAÇÃO NO REAL VALOR DA VARIÁVEL, VC PRECISA PASSAR O PONTEIRO,
	// NO CASO PRA SEU CÓDIGO INTEIRO
	// QUANDO VC SÓ PRECISA FAZER UMA CÓPIA, NÃO PRECISA
	fmt.Println("Funções com Ponteiro")

	// INVERTENDO COM PONTEIRO
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//INVERTENDO SEM PONTEIROS
	novoNumero := 40
	fmt.Println(novoNumero)
	// AQUI TO PASSANDO O ENDEREÇO PARA A FUNÇÃO
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
