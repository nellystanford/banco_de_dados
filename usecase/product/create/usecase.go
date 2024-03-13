package create

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db/product"
)

func CreateItem(c *fiber.Ctx, db *sql.DB, input Input) (Output, error) {
	if !inputIsValid(input) {
		return Output{}, fmt.Errorf("Input is invalid")
	}

	item, err := database.Create(c, db, buildInput(input))
	if err != nil {
		return Output{}, err
	}

	return buildOutput(item), nil
}

func inputIsValid(input Input) bool {
	if input.Name == "" || input.Color == "" || input.Code == 0 || input.Size == 0 || input.Value == 0 {
		return false
	}
	return true
}

func buildInput(input Input) entity.Product {
	return entity.Product{
		Name:     input.Name,
		Size:     input.Size,
		Color:    input.Color,
		Quantity: input.Quantity,
		Code:     input.Code,
		Value:    input.Value,
	}
}

func buildOutput(entity entity.Product) Output {
	return Output{
		ID:       entity.ID,
		Name:     entity.Name,
		Size:     entity.Size,
		Color:    entity.Color,
		Quantity: entity.Quantity,
		Code:     entity.Code,
		Value:    entity.Value,
	}
}
