package update

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db"
)

func UpdateItem(c *fiber.Ctx, db *sql.DB, input Input) (Output, error) {
	if !inputIsValid(input) {
		return Output{}, fmt.Errorf("Input is invalid")
	}

	product, err := database.FindByNameColorAndSize(c, db, input.Nome, input.Cor, input.Tamanho)
	if err != nil {
		return Output{}, err
	}
	if product.ID == 0 {
		return Output{}, fmt.Errorf("item does not exist in database")
	}

	if input.Nome != product.Nome ||
		input.Tamanho != product.Tamanho ||
		input.Cor != product.Cor ||
		input.Codigo != product.Codigo {
		return Output{}, fmt.Errorf("unable to modify name, size, color or code")
	}

	product.Valor = input.Valor
	product.Quantidade = input.Quantidade
	item, err := database.Update(c, db, product)
	if err != nil {
		return Output{}, err
	}

	return buildOutput(item), nil
}

func inputIsValid(input Input) bool {
	if input.Quantidade == 0 || input.Valor == 0 {
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
