package http

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	createProduct "github.com/nellystanford/banco_de_dados/usecase/product/create"
	deleteProduct "github.com/nellystanford/banco_de_dados/usecase/product/delete"
	findProduct "github.com/nellystanford/banco_de_dados/usecase/product/find"
	findProductByName "github.com/nellystanford/banco_de_dados/usecase/product/find_by_name"
	findOneProduct "github.com/nellystanford/banco_de_dados/usecase/product/find_one"
	updateProduct "github.com/nellystanford/banco_de_dados/usecase/product/update"

	createClient "github.com/nellystanford/banco_de_dados/usecase/client/create"
	deleteClient "github.com/nellystanford/banco_de_dados/usecase/client/delete"
	findClient "github.com/nellystanford/banco_de_dados/usecase/client/find"

	createOrder "github.com/nellystanford/banco_de_dados/usecase/order/create"
	deleteOrder "github.com/nellystanford/banco_de_dados/usecase/order/delete"
	findOrder "github.com/nellystanford/banco_de_dados/usecase/order/find"

	"github.com/nellystanford/banco_de_dados/usecase/summary"
)

func Routes(db *sql.DB) {
	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/findAllProducts", func(c *fiber.Ctx) error {
		return findProduct.Handler(c, db)
	})

	app.Get("/findOneProduct", func(c *fiber.Ctx) error {
		return findOneProduct.Handler(c, db)
	})

	app.Get("/findProductByName", func(c *fiber.Ctx) error {
		return findProductByName.Handler(c, db)
	})

	app.Post("/createProduct", func(c *fiber.Ctx) error {
		return createProduct.Handler(c, db)
	})

	app.Patch("/updateProduct", func(c *fiber.Ctx) error {
		return updateProduct.Handler(c, db)
	})

	app.Delete("/deleteProduct", func(c *fiber.Ctx) error {
		return deleteProduct.Handler(c, db)
	})

	app.Get("/findClient", func(c *fiber.Ctx) error {
		return findClient.Handler(c, db)
	})

	app.Post("/createClient", func(c *fiber.Ctx) error {
		return createClient.Handler(c, db)
	})

	app.Delete("/deleteClient", func(c *fiber.Ctx) error {
		return deleteClient.Handler(c, db)
	})

	app.Get("/findOrder", func(c *fiber.Ctx) error {
		return findOrder.Handler(c, db)
	})

	app.Post("/createOrder", func(c *fiber.Ctx) error {
		return createOrder.Handler(c, db)
	})

	app.Delete("/deleteOrder", func(c *fiber.Ctx) error {
		return deleteOrder.Handler(c, db)
	})

	app.Get("/summary", func(c *fiber.Ctx) error {
		return summary.Handler(c, db)
	})

	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
