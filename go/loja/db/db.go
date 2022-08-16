package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Println("Erro ao conectar ao banco de dados", err)
	}

	return db
}
