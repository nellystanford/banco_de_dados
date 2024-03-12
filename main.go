package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nellystanford/banco_de_dados/http"
)

// adicionar verificação de existencia de item na tabela ao criar no usecase
// modificar nomes do entity para ingles

// adicionar função de gerar relatório
// fazer modelagem das classes que serão utilizadas no sistema utilizando diagrama UML de classe
// adicionar quantidade vendia, fazer relatorio getsummary e relatar valor total ventido, quantidade de cada peça restante e 3 peças mais populares

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
