package main

import "fmt"

func main() {
	// SIMILAR AS STRUCTS, PORÉM O MAP TEM UMA CHAVE E UM VALOR TODOS DO MESMO TIPO.
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
	}
	fmt.Println(usuario2)
}
