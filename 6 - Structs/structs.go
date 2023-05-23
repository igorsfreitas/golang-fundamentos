package main

import "fmt"

// STRUCTS EM GO SÃO TIPOS DE DADOS PERSONALIZADOS - UMA ESTRUTURA QUE FUNCIONA COMO UM OBJETO / CLASSE
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	var u usuario
	fmt.Println(u)

	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Davi", 21, enderecoExemplo} // FORMA DE DECLARAR STRUCTS POR INFERENCIA DE TIPO
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"} // FORMA DE DECLARAR STRUCTS POR INFERENCIA DE TIPO EM UM CAMPO ESPECIFICO
	fmt.Println(usuario3)
}
