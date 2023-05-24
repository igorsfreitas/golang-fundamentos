package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Groups")

	var waitgroup sync.WaitGroup

	waitgroup.Add(2) // AQUI DIZ QUE ESSE WAITGROUP ESPERA 2 GO ROUTINES

	go func() {
		escrever("Olá Mundo")
		waitgroup.Done() // AQUI DIZ QUE ESSA GO ROUTINE TERMINOU ENTÃO SUBTRAI 1 DO WAITGROUP
	}()

	go func() {
		escrever("Programando em Go!")
		waitgroup.Done()
	}()

	waitgroup.Wait() // AQUI DIZ QUE O PROGRAMA ESPERA TODAS AS GO ROUTINES TERMINAREM

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
