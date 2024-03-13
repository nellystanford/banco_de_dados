package create

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
	database "github.com/nellystanford/banco_de_dados/infra/db/client"
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
	if input.Name == "" || input.Email == "" || input.CPF == "" || input.Address == "" {
		return false
	}
	return true
}

func buildInput(input Input) entity.Client {
	return entity.Client{
		Name:       input.Name,
		Email:      input.Email,
		CPF:        input.CPF,
		Address:    input.Address,
		Newsletter: input.Newsletter,
	}
}

func buildOutput(entity entity.Client) Output {
	return Output{
		ID:         entity.ID,
		Name:       entity.Name,
		Email:      entity.Email,
		CPF:        entity.CPF,
		Address:    entity.Address,
		Newsletter: entity.Newsletter,
	}
}
