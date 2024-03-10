package delete

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	database "github.com/nellystanford/banco_de_dados/infra/db"
)

func DeleteItem(c *fiber.Ctx, db *sql.DB, input Input) error {
	err := database.Delete(c, db, input.ID)
	if err != nil {
		return err
	}

	return nil
}
