package find

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db/product"
)

func FindAll(c *fiber.Ctx, db *sql.DB) ([]Output, error) {
	items, err := database.Find(c, db)
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
		Name:     entity.Name,
		Size:     entity.Size,
		Color:    entity.Color,
		Quantity: entity.Quantity,
		Code:     entity.Code,
		Value:    entity.Value,
	}
}
