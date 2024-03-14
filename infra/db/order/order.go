package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/nellystanford/banco_de_dados/entity"
)

type ClientDBModel struct {
	ID         int
	ClientName string
	CPF        string
	Address    string
	Product    pq.Int32Array
}

func Create(c *fiber.Ctx, db *sql.DB, o entity.Order) error {
	fmt.Printf("Inserting order into database...\n")

	_, err := db.Query("INSERT INTO pedidos (nome, cpf, endereco, produtos) VALUES ($1, $2, $3, $4)", o.ClientName, o.CPF, o.Address, o.Product)
	if err != nil {
		log.Fatalf("An error occured while executing insertion: %v", err)
		return err
	}

	return nil
}

func Delete(c *fiber.Ctx, db *sql.DB, id int) error {
	fmt.Printf("Deleting order from database...\n")

	_, err := db.Query("DELETE FROM pedidos WHERE id = $1", id)
	if err != nil {
		log.Fatalf("An error occured while executing deletion: %v", err)
		return err
	}

	return nil
}

func Find(c *fiber.Ctx, db *sql.DB) ([]entity.Order, error) {
	fmt.Printf("Finding all orders...\n")

	ord, err := db.Query("SELECT * FROM pedidos")
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return []entity.Order{}, err
	}

	orders, err := sqlResultToEntity(ord)
	if err != nil {
		log.Fatalf("An error occured while transcripting valued recovered from table: %v", err)
		return []entity.Order{}, err
	}

	return orders, nil
}

func sqlResultToEntity(p *sql.Rows) ([]entity.Order, error) {
	var order entity.Order
	var orders []entity.Order
	var i int = 0

	for p.Next() {
		var name, cpf, address string
		var product pq.Int32Array
		var id int

		err := p.Scan(&id, &name, &cpf, &address, &product)
		if err != nil {
			return []entity.Order{}, err
		}
		order.ID = id
		order.ClientName = name
		order.CPF = cpf
		order.Address = address
		order.Product = product

		orders = append(orders, order)
		i++
	}

	return orders, nil
}
