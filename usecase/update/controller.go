package update

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	input := Input{}

	if err := c.BodyParser(&input); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	out, err := UpdateItem(c, db, input)
	if err != nil {
		return err
	}

	return c.JSON(out)
}
