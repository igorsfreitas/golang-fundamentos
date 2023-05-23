package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoEstaAprovado(n1, n2 int) bool {
	defer fmt.Println("Média calculada, o valor será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media > 6 {

		return true
	}

	return false
}

func main() {
	// ESSA FUNÇÃO É MUITO UTILIZADA PARA FECHAR
	//A CONEXÃO AO BANCO DE DADOS, POIS ELA SERÁ EXECUTADA NO FINAL DO PROGRAMA,
	//ANTES DO RETURN
	fmt.Println(alunoEstaAprovado(7, 8))
}
