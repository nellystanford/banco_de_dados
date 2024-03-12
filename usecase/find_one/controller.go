package find_one

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	var input Input

	input.Nome = c.Query("name")
	input.Cor = c.Query("color")
	tamanho, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return err
	}
	input.Tamanho = tamanho

	out, err := FindOne(c, db, input)
	if err != nil {
		return err
	}

	return c.JSON(out)
}
