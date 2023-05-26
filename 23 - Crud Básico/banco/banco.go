package banco

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver do Postgres
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	connStr := "postgres://postgres:postgres@localhost/devbook-localhost?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
