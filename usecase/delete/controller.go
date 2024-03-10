package delete

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	input := Input{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return err
	}

	input.ID = id
	err = DeleteItem(c, db, input)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}
