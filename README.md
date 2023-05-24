1 - Toda função inicial deve estar dentro de um pacote main
2 - As funções privadas, são escritas com letra minuscula, e as públicas com letra maiuscula (Escrever - Está publica e exportada, escrever - Está privada)
3 - Para instalar um pacote externo: go get URL_DO_PACOTE
4 - Para buildar: go build
5 - Para rodar: go run NOME_DO_ARQUIVO
6 - Para criar módulo: go mod init NOME_DO_MODULO
7 - Para remover dependencias nao utilizadas: go mod tidy
8 - O go não é uma linguagem orientada a objetos, não tem classes, não tem polimorfismo, não tem herança
9 - Para rodar os testes: go test ./...
10 - Para rodar os testes no modo verboso: go test -v ./...
11 - Para verificar a cobertura de testes: go test --cover
12 - Para gerar relatório com a cobertura de testes: go test --coverprofile resultado.txt ./...
13 - Para ler o relatório gerado e jogar no terminal: go tool cover --func=resultado.txt
14 - Gera um html do relatório gerado: go tool cover --html=resultado.txt

DUVIDAS

1 - Como fazer injecao / inversão de dependencias no GO se não tem construtor
2 - Como fazer testes unitários
