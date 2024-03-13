package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/nellystanford/banco_de_dados/entity"
)

type ProductDBModel struct {
	ID         int
	Nome       string
	Tamanho    int
	Cor        string
	Quantidade int
	Codigo     int
	Valor      float64
}

func Create(c *fiber.Ctx, db *sql.DB, p entity.Product) (entity.Product, error) {
	fmt.Printf("Inserting product into database...\n")

	_, err := db.Query("INSERT INTO produtos (nome, tamanho, cor, quantidade, codigo, valor) VALUES ($1, $2, $3, $4, $5, $6)", p.Name, p.Size, p.Color, p.Quantity, p.Code, p.Value)
	if err != nil {
		log.Fatalf("An error occured while executing insertion: %v", err)
		return entity.Product{}, err
	}

	product, err := FindByNameColorAndSize(c, db, p.Name, p.Color, p.Size)
	if err != nil {
		log.Fatalf("An error occured retrieving data after insertion: %v", err)
		return entity.Product{}, err
	}

	return product, nil
}

func Update(c *fiber.Ctx, db *sql.DB, product entity.Product) (entity.Product, error) {
	fmt.Printf("Updating item...\n")

	_, err := db.Query("UPDATE produtos SET quantidade = $1, valor = $2 WHERE id = $3", product.Quantity, product.Value, product.ID)
	if err != nil {
		log.Fatalf("An error occured while executing insertion: %v", err)
		return entity.Product{}, err
	}

	return product, nil
}

func Delete(c *fiber.Ctx, db *sql.DB, id int) error {
	fmt.Printf("Deleting product from database...\n")

	_, err := db.Query("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		log.Fatalf("An error occured while executing deletion: %v", err)
		return err
	}

	return nil
}

func Find(c *fiber.Ctx, db *sql.DB) ([]entity.Product, error) {
	fmt.Printf("Finding all items...\n")

	p, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return []entity.Product{}, err
	}

	products, err := sqlResultToEntity(p)
	if err != nil {
		log.Fatalf("An error occured while transcripting valued recovered from table: %v", err)
		return []entity.Product{}, err
	}

	return products, nil
}

func FindByName(c *fiber.Ctx, db *sql.DB, name string) ([]entity.Product, error) {
	fmt.Printf("Finding all items by name...\n")

	p, err := db.Query("SELECT * FROM produtos WHERE nome = $1", name)
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return []entity.Product{}, err
	}

	products, err := sqlResultToEntity(p)
	if err != nil {
		log.Fatalf("An error occured while transcripting valued recovered from table: %v", err)
		return []entity.Product{}, err
	}

	return products, nil
}

func FindByNameColorAndSize(c *fiber.Ctx, db *sql.DB, name string, color string, size int) (entity.Product, error) {
	fmt.Printf("Finding by name, color and size...\n")

	p, err := db.Query("SELECT * FROM produtos WHERE nome = $1 and cor = $2 and tamanho = $3", name, color, size)
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return entity.Product{}, err
	}

	products, err := sqlResultToEntity(p)
	if err != nil {
		log.Fatalf("An error occured while transcripting valued recovered from table: %v", err)
		return entity.Product{}, err
	}
	if len(products) == 0 {
		return entity.Product{}, nil
	}

	return products[0], nil
}

func sqlResultToEntity(p *sql.Rows) ([]entity.Product, error) {
	var product entity.Product
	var products []entity.Product
	var i int = 0

	for p.Next() {
		var name, color string
		var id, size, quantity, code int
		var value float64

		err := p.Scan(&id, &name, &size, &color, &quantity, &code, &value)
		if err != nil {
			return []entity.Product{}, err
		}
		product.ID = id
		product.Name = name
		product.Size = size
		product.Color = color
		product.Quantity = quantity
		product.Code = code
		product.Value = value

		products = append(products, product)
		i++
	}

	return products, nil
}
