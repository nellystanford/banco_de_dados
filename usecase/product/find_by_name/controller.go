package find_by_name

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	var input Input
	input.Nome = c.Query("nome")

	out, err := Find(c, db, input)
	if err != nil {
		return err
	}

	return c.JSON(out)
}
