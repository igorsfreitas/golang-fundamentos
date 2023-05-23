package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// ESTA SINTAX DIZ QUE TODOS OS USUÁRIOS
// TEM O MÉTODO SALVAR
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s\n", u.nome)
}

// PARA EXECUTAR AÇÕES, OU FAZER VALIDAÇÃO
// SEM ALTERAR OS DADOS DO STRUCT
// NÃO DEVE USAR PONTEIRO
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// PARA ALTERAR OS DADOS DESSE STRUCT, EU DEVO
// UTILIZAR PONTEIRO PARA PASSAR A REFERENCIA DO STRUCT
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	//MÉTODOS SÃO UTILIZADOS PARA AUXILIAREM AS STRUCTS

	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.maiorDeIdade()

	usuario2 := usuario{"Davi", 2}
	usuario2.salvar()
	usuario2.maiorDeIdade()

	usuario2.fazerAniversario()
	fmt.Println(usuario2)
}
