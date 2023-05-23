package main

import "fmt"

func main() {

	func() {
		fmt.Println("Funções Anonimas")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Funções Anonimas -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
