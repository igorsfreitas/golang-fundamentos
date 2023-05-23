package main

import "fmt"

func main() {
	// Tipos números inteiros: int8, int16, int32, int64,
	// Tipos números inteiros positivos: uint8, uint16, uint32, uint64

	var numero int16 = 100
	fmt.Println(numero)

	// int32 = rune
	var numero2 rune = 123456

	fmt.Println(numero2)

	// uint8 = byte
	var numero3 byte = 123

	fmt.Println(numero3)

	// Tipos números reais: float32, float64

	var numeroReal1 float32 = 123.45

	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45

	fmt.Println(numeroReal2)

	str2 := "Texto"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//FIM STRINGS

	// Tipos default do go para numeros é 0, para strings é vazio, boolean é false e para erro é nil (nulo)
}
