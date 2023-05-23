package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interfaces Tipos genéricos")

	generica("String")
	generica(1)
	generica(true)

	fmt.Println("Mapas", 1, true) // EXEMPLO DE FUNCAO QUE USA FUNCAO GENERICA
}
