package entity

import "github.com/lib/pq"

type Order struct {
	ID         int
	ClientName string
	CPF        string
	Address    string
	Product    pq.Int32Array
}
