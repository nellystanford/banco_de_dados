package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nellystanford/banco_de_dados/entity"
)

type ClientDBModel struct {
	ID         int
	Name       string
	Email      string
	CPF        string
	Address    string
	Newsletter bool
}

func Create(c *fiber.Ctx, db *sql.DB, p entity.Client) (entity.Client, error) {
	fmt.Printf("Inserting client into database...\n")

	_, err := db.Query("INSERT INTO clientes (nome, email, cpf, endereco, newsletter) VALUES ($1, $2, $3, $4, $5)", p.Name, p.Email, p.CPF, p.Address, p.Newsletter)
	if err != nil {
		log.Fatalf("An error occured while executing insertion: %v", err)
		return entity.Client{}, err
	}

	client, err := FindByCPF(c, db, p.CPF)
	if err != nil {
		log.Fatalf("An error occured retrieving data after insertion: %v", err)
		return entity.Client{}, err
	}

	return *client, nil
}

func Delete(c *fiber.Ctx, db *sql.DB, id int) error {
	fmt.Printf("Deleting client from database...\n")

	_, err := db.Query("DELETE FROM clientes WHERE id = $1", id)
	if err != nil {
		log.Fatalf("An error occured while executing deletion: %v", err)
		return err
	}

	return nil
}

func Find(c *fiber.Ctx, db *sql.DB) ([]entity.Client, error) {
	fmt.Printf("Finding all clients...\n")

	ct, err := db.Query("SELECT * FROM clientes")
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return []entity.Client{}, err
	}

	clients, err := sqlResultToEntity(ct)
	if err != nil {
		log.Fatalf("An error occured while transcripting valued recovered from table: %v", err)
		return []entity.Client{}, err
	}

	return clients, nil
}

func FindByCPF(c *fiber.Ctx, db *sql.DB, cpf string) (*entity.Client, error) {
	fmt.Printf("Finding client by cpf...\n")

	p, err := db.Query("SELECT * FROM clientes WHERE cpf = $1", cpf)
	if err != nil {
		log.Fatalf("An error occured while recovering data from table: %v", err)
		return nil, err
	}

	clients, err := sqlResultToEntity(p)
	if err != nil {
		log.Fatalf("An error occured while transcripting valued recovered from table: %v", err)
		return nil, err
	}
	if len(clients) == 0 {
		return nil, nil
	}

	return &clients[0], nil
}

func sqlResultToEntity(p *sql.Rows) ([]entity.Client, error) {
	var client entity.Client
	var clients []entity.Client
	var i int = 0

	for p.Next() {
		var name, email, cpf, address string
		var newsletter bool
		var id int

		err := p.Scan(&id, &name, &email, &cpf, &address, &newsletter)
		if err != nil {
			return []entity.Client{}, err
		}
		client.ID = id
		client.Name = name
		client.Email = email
		client.CPF = cpf
		client.Address = address
		client.Newsletter = newsletter

		clients = append(clients, client)
		i++
	}

	return clients, nil
}
