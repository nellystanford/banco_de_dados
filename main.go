package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nellystanford/banco_de_dados/http"
)

// verificacoes necessarias: se cpf ja existe ao criar cliente e se produto ja existe
// atualizar diagrama uml

func main() {
	db := connectWithDatabase()
	http.Routes(db)
}

func connectWithDatabase() *sql.DB {
	connection := "user=postgres dbname=ns_loja password=210110 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
