package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w.Write([]byte("Olá Mundo!")))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
	fmt.Println("Servidor rodando na porta 5000")
}
