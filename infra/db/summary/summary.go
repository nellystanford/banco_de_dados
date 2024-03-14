package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetOrdersAmount(c *fiber.Ctx, db *sql.DB) (int, error) {
	fmt.Printf("Getting orders amount...\n")

	orders, err := db.Query("SELECT COUNT(*) FROM pedidos")
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return 0, err
	}

	var count int
	for orders.Next() {
		if err := orders.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count, nil
}

func GetClientsAmount(c *fiber.Ctx, db *sql.DB) (int, error) {
	fmt.Printf("Getting orders amount...\n")

	sales, err := db.Query("SELECT COUNT (*) FROM clientes")
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return 0, err
	}

	var count int
	for sales.Next() {
		if err := sales.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count, nil
}

func GetNewsletterSubscriptions(c *fiber.Ctx, db *sql.DB) (int, error) {
	fmt.Printf("Getting orders amount...\n")

	newsletter, err := db.Query("SELECT COUNT (*) FROM clientes WHERE newsletter = $1", true)
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return 0, err
	}

	var count int
	for newsletter.Next() {
		if err := newsletter.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count, nil
}

// func GetSalesAmount(c *fiber.Ctx, db *sql.DB) (float64, error) {
// 	fmt.Printf("Getting orders amount...\n")

// 	orders, err := db.Query("SELECT COUNT (*) FROM pedidos")
// 	if err != nil {
// 		log.Fatalf("An error occured while recovering data from table: %v", err)
// 		return 0, err
// 	}

// 	return 0, nil
// }
