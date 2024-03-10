package find

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db"
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
		ID:         entity.ID,
		Nome:       entity.Nome,
		Tamanho:    entity.Tamanho,
		Cor:        entity.Cor,
		Quantidade: entity.Quantidade,
		Codigo:     entity.Codigo,
		Valor:      entity.Valor,
	}
}
