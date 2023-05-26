package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Driver do Postgres
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/devbook-localhost?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado com sucesso!")
}
