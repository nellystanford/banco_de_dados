package http

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/usecase/create"
	"github.com/nellystanford/banco_de_dados/usecase/delete"
	"github.com/nellystanford/banco_de_dados/usecase/find"
	"github.com/nellystanford/banco_de_dados/usecase/find_by_name"
	"github.com/nellystanford/banco_de_dados/usecase/find_one"
	"github.com/nellystanford/banco_de_dados/usecase/update"
)

func Routes(db *sql.DB) {
	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/find", func(c *fiber.Ctx) error {
		return find.Handler(c, db)
	})

	app.Get("/find-one", func(c *fiber.Ctx) error {
		return find_one.Handler(c, db)
	})

	app.Get("/find-by-name", func(c *fiber.Ctx) error {
		return find_by_name.Handler(c, db)
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
