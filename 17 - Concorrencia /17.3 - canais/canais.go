package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Canais")

	canal := make(chan string) // 1 - CRIANDO UM CANAL

	go escrever("Olá Mundo", canal) // 2 - ADICIONANDO GO ROUTINE AO CANAL

	// for {
	// 	mensagem, aberto := <-canal // 4 - RECEBENDO DADOS DO CANAL
	// 	if !aberto {                // 6 - AQUI ESTOU VERIFICANDO SE O CANAL ESTÁ ABERTO OU FECHADO
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	} // 7 - OUTRA FORMA DE RECEBER DADOS DO CANAL MAIS RESUMIDA
	// SEM PRECISAR VERIFICAR SE O CANAL ESTÁ ABERTO OU FECHADO E SEM O BREAK

	fmt.Println("Fim do programa")

	// UM CANAL É UM CANAL DE COMUNICAÇÃO , QUE PODE TANTO ENVIAR COMO RECEBER DADOS,
	// ONDE USAMOS ESSES DADOS PARA SINCRONIZAR AS NOSSAS GO ROUTINES

	// ATENÇÃO SEMPRE PARA VERIFICAR SE O CANAL ESTÁ ABERTO OU FECHADO,
	// POIS SE NÃO ESTIVER ABERTO E CONTINUAR ESPERANDO ALGUMA MENSAGEM DESSE CANAL
	// PODE DAR UM DEADLOCK, E ISSO NÃO VAI EM TEMPO DE COMPILAÇÃO, MAS SIM EM TEMPO DE EXECUÇÃO
	// ENTÃO PODE CHEGAR A DAR ERROS EM PRODUÇÃO.

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // 3 - ENVIANDO DADOS PARA O CANAL
		time.Sleep(time.Second)
	}

	close(canal) // 5 - AQUI INFORMA QUE ESTOU FECHANDO O CANAL
}
