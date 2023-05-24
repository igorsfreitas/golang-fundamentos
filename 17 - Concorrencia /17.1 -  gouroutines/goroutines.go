package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Routines")
	// CONCORRENCIA != PARALELISMO
	go escrever("Olá Mundo") //GO ROUTINEs
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
