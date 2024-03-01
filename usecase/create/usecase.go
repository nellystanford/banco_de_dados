package create

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	database "github.com/nellystanford/banco_de_dados/infra/db"
)

func CreateItem(c *fiber.Ctx, db *sql.DB, input Input) error {
	if !inputIsValid(input) {
		return fmt.Errorf("Input is invalid")
	}

	fmt.Printf("Received input %v", input)
	err := database.Create(c, db, database.ProductDBModel(input))
	if err != nil {
		return err
	}

	return nil
}

func inputIsValid(input Input) bool {
	if input.Nome == "" || input.Cor == "" || input.Código == 0 || input.Tamanho == 0 || input.Valor == 0 {
		return false
	}
	return true
}
