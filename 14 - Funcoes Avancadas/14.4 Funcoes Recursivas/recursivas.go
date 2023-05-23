package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	} else {
		return fibonacci(posicao-2) + fibonacci(posicao-1)
	}
}

func main() {
	// FUNCOES RECURSIVAS SÃO FUNÇÕES QUE CHAMAM ELAS MESMAS
	// AS FUNCOES RECURSIVAS NÃO SÃO RECOMENDADAS PARA SEREM USADAS EM CASOS MUITO GRANDES, POIS PODEM CAUSAR ESTOURO DE PILHA (OVERFLOW)
	fmt.Println("Funções Recursivas")

	posicao := uint(1)
	fmt.Println(fibonacci(posicao))

}
