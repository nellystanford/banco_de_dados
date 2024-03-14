package create

import "github.com/lib/pq"

type Input struct {
	ClientName string        `json:"nome_do_cliente"`
	CPF        string        `json:"cpf"`
	Address    string        `json:"endereco"`
	Product    pq.Int32Array `json:"produtos"`
}

type Output struct {
	ClientName string        `json:"nome_do_cliente"`
	CPF        string        `json:"cpf"`
	Address    string        `json:"endereco"`
	Product    pq.Int32Array `json:"produtos"`
}
