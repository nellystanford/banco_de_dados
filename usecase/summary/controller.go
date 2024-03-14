package summary

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	out, err := GetSummary(c, db)
	if err != nil {
		return err
	}

	return c.JSON(out)
}
