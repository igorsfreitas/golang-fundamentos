package main

import (
	"fmt"
)

func main() {
	fmt.Println("Canais com buffer")

	canal := make(chan string, 2) // A DIFERENÇA DO CANAL COM BUFFER PARA O CANAL NORMAL
	//	É QUE O CANAL COM BUFFER TEM UM SEGUNDO PARAMETRO QUE É A
	// QUANTIDADE DE MENSAGENS QUE ELE VAI ARMAZENAR, NO CASO 2

	canal <- "Olá Mundo"
	canal <- "Programando em Go!"
	//canal <- "DeadLock" // NÃO VAI FUNCIONAR POIS O CANAL SÓ TEM CAPACIDADE PARA 2 MENSAGENS

	mensagem := <-canal
	fmt.Println(mensagem)
}
