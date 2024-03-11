package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/usecase/create"
	"github.com/nellystanford/banco_de_dados/usecase/delete"
	"github.com/nellystanford/banco_de_dados/usecase/find"
	"github.com/nellystanford/banco_de_dados/usecase/update"
)

// adicionar verificação de existencia de item na tabela ao criar no usecase
// adicionar rota de achar um item e de achar por nome
// modificar nomes para ingles
// adicionar função de gerar relatório
// fazer modelagem das classes que serão utilizadas no sistema utilizando diagrama UML de classe

func main() {
	db := connectWithDatabase()

	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/find", func(c *fiber.Ctx) error {
		return find.Handler(c, db)
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		return create.Handler(c, db)
	})

	app.Patch("/update", func(c *fiber.Ctx) error {
		return update.Handler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return delete.Handler(c, db)
	})

	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func connectWithDatabase() *sql.DB {
	connection := "user=postgres dbname=ns_loja password=210110 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
