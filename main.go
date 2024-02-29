package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=ns_loja password=210110 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	fmt.Println("Opening db connection...")
	db := conectaComBancoDeDados()

	fmt.Println("Closing db connection...")
	defer db.Close()
}
