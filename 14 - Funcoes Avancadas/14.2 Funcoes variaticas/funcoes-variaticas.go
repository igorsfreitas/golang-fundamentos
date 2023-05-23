package main

import "fmt"

func soma(numeros ...int) (total int) { // RECEBE UMA QUANTIDADE INDETERMINADA DE PARAMETROS (SLICE)

	for _, numero := range numeros {
		total += numero
	}

	return
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variaticas")

	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	escrever("Olá Mundo", 1, 2)
}
