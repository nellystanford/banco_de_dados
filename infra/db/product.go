package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type ProductDBModel struct {
	Nome       string
	Tamanho    int
	Cor        string
	Quantidade int
	Código     int
	Valor      float64
}

func Create(c *fiber.Ctx, db *sql.DB, p ProductDBModel) error {
	fmt.Printf("Inserting product into database...")

	_, err := db.Exec("INSERT INTO produtos (nome, tamanho, cor, quantidade, codigo, valor) VALUES ($1, $2, $3, $4, $5, $6)", p.Nome, p.Tamanho, p.Cor, p.Quantidade, p.Código, p.Valor)
	if err != nil {
		log.Fatalf("An error occured while executing insertion: %v", err)
		return err
	}

	return nil
}
