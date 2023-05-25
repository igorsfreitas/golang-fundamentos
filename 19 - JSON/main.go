package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON)                  // QUANDO ELE É CONVERTIDO PARA JSON, O STRUCT VIRA UM ARRAY DE BYTES
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // DESSA FORMA, CONVERTE-SE PARA JSON
}
