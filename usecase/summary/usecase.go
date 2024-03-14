package summary

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	database "github.com/nellystanford/banco_de_dados/infra/db/summary"
)

func GetSummary(c *fiber.Ctx, db *sql.DB) (Output, error) {
	var output Output

	// total orders
	totalOrders, err := database.GetOrdersAmount(c, db)
	if err != nil {
		return Output{}, err
	}
	output.TotalOrders = totalOrders

	// total clients
	TotalClients, err := database.GetClientsAmount(c, db)
	if err != nil {
		return Output{}, err
	}
	output.TotalClients = TotalClients

	// total newsletter subscriptions
	NewsletterSubscriptions, err := database.GetNewsletterSubscriptions(c, db)
	if err != nil {
		return Output{}, err
	}
	output.NewsletterSubscriptions = NewsletterSubscriptions

	// sales amount
	SalesAmount, err := database.GetSalesAmount(c, db)
	if err != nil {
		return Output{}, err
	}
	output.SalesAmount = SalesAmount

	return output, nil
}
