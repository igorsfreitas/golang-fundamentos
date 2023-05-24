package main

import (
	"fmt"
)

func main() {
	fmt.Println("Worker Pools")

	// WORKER POOL É UMA FILA DE TAREFAS A SEREM EXECUTADAS,
	// PEGANDO ITENS DESSA FILA DE MANEIRA INDEPENDENTE

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for numero := 0; numero < 45; numero++ {
		tarefas <- numero
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
