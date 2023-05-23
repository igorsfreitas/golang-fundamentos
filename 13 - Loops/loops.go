package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		// time.Sleep(time.Second)
		fmt.Println(i)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes { // AQUI É COMO SE FOSSE UM FOREACH EX: arr.forEach((item, index) => {})
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)         // DESSA FORMA POR PADRÃO ELE VAI TRAZER O INDICE E O VALOR DA TABELA ASC DO CHAR DA PALAVRA
		fmt.Println(indice, string(letra)) // DESSA FORMA ELE VAI TRAZER O INDICE E O VALOR DA LETRA
	}

	// NÃO SE PODE UTILIZAR LOOPS EM STRUCT, APENAS EM MAPS
}
