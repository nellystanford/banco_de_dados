package create

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db/order"
)

func CreateItem(c *fiber.Ctx, db *sql.DB, input Input) (Output, error) {
	if !inputIsValid(input) {
		return Output{}, fmt.Errorf("Input is invalid")
	}

	err := database.Create(c, db, buildInput(input))
	if err != nil {
		return Output{}, err
	}

	return Output(input), nil
}

func inputIsValid(input Input) bool {
	if input.ClientName == "" || input.CPF == "" || input.Address == "" || len(input.Product) == 0 {
		return false
	}
	return true
}

func buildInput(input Input) entity.Order {
	return entity.Order{
		ClientName: input.ClientName,
		CPF:        input.CPF,
		Address:    input.Address,
		Product:    input.Product,
	}
}
