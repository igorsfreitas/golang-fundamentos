package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select")
	// O SELECT TEM A SINTAXE PARECIDA COM FOR, MAS É UTILIZADO PARA CONSUMIR DADOS DE CANAIS

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}
}
