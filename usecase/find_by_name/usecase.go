package find_by_name

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db"
)

func Find(c *fiber.Ctx, db *sql.DB, input Input) ([]Output, error) {
	items, err := database.FindByName(c, db, input.Nome)
	if err != nil {
		return []Output{}, err
	}

	var output []Output
	for _, item := range items {
		output = append(output, buildOutput(item))
	}

	return output, nil
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
