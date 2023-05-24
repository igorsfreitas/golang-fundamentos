package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavra := strings.Split(endereco, " ")[0]

	//verificar se a primeira palavra do endereço é valida

	for _, tipo := range tiposValidos {
		if tipo == strings.ToLower(primeiraPalavra) {
			return primeiraPalavra
		}
	}

	return "Tipo Inválido"
}
