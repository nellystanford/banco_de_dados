package create

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db"
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
		Nome:       input.Name,
		Tamanho:    input.Size,
		Cor:        input.Color,
		Quantidade: input.Quantity,
		Codigo:     input.Code,
		Valor:      input.Value,
	}
}

func buildOutput(entity entity.Product) Output {
	return Output{
		ID:       entity.ID,
		Name:     entity.Nome,
		Size:     entity.Tamanho,
		Color:    entity.Cor,
		Quantity: entity.Quantidade,
		Code:     entity.Codigo,
		Value:    entity.Valor,
	}
}
