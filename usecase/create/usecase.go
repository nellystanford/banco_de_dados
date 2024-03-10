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
	if input.Nome == "" || input.Cor == "" || input.Codigo == 0 || input.Tamanho == 0 || input.Valor == 0 {
		return false
	}
	return true
}

func buildInput(input Input) entity.Product {
	return entity.Product{
		Nome:       input.Nome,
		Tamanho:    input.Tamanho,
		Cor:        input.Cor,
		Quantidade: input.Quantidade,
		Codigo:     input.Codigo,
		Valor:      input.Valor,
	}
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
