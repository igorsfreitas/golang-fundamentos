package main

import "fmt"

func clojure() func() {
	texto := "Dentro da função clojure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	fmt.Println("Funções Clojure")
	// SÃO FUNÇÕES QUE REFERENCIAM VARIÁVEIS QUE RETORNAM OUTRA FUNÇÃO

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := clojure()
	funcaoNova()
}
