package main

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	resultado := somar(10, 20)
	println(resultado)

	// CASO NAO PRECISE DO VALOR DE RETORNO, PODE-SE USAR O UNDERLINE PARA IGNORAR O RETORNO
	// EX: _, subtracao := calculosMatematicos(10, 15)
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	println(resultadoSoma, resultadoSubtracao)
}
