package find_one

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db"
)

func FindOne(c *fiber.Ctx, db *sql.DB, input Input) (Output, error) {
	item, err := database.FindByNameColorAndSize(c, db, input.Nome, input.Cor, input.Tamanho)
	if err != nil {
		return Output{}, err
	}

	return buildOutput(item), nil
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
