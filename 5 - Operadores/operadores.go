package main

import "fmt"

func main() {
	// OPERADORES ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	soma := numero1 + numero2 // ESSA LINHA DÁ ERRADO PQ NÃO PODE SOMAR TIPOS DIFERENTES
	fmt.Println(soma)

	// FIM DOS ARITMETICOS

	// OPERADORES DE ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNARIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)
	numero *= 3 // numero = numero * 3
	fmt.Println(numero)
	numero /= 10 // numero = numero / 10
	fmt.Println(numero)
	numero %= 3 // numero = numero % 3
	fmt.Println(numero)
	// FIM DOS OPERADORES UNARIOS

	// OPERADORES TERNARIOS NÃO EXISTEM EM GO
}
