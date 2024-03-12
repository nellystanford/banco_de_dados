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

	product, err := database.FindByNameColorAndSize(c, db, input.Name, input.Color, input.Size)
	if err != nil {
		return Output{}, err
	}
	if product.ID == 0 {
		return Output{}, fmt.Errorf("item does not exist in database")
	}

	if input.Name != product.Nome ||
		input.Size != product.Tamanho ||
		input.Color != product.Cor ||
		input.Code != product.Codigo {
		return Output{}, fmt.Errorf("unable to modify name, size, color or code")
	}

	product.Valor = input.Value
	product.Quantidade = input.Quantity
	item, err := database.Update(c, db, product)
	if err != nil {
		return Output{}, err
	}

	return buildOutput(item), nil
}

func inputIsValid(input Input) bool {
	if input.Quantity == 0 || input.Value == 0 {
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
