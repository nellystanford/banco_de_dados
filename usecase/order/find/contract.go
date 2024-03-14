package find

import "github.com/lib/pq"

type Output struct {
	ID         int           `json:"id"`
	ClientName string        `json:"nome_do_cliente"`
	CPF        string        `json:"cpf"`
	Address    string        `json:"endereco"`
	Product    pq.Int32Array `json:"produtos"`
}
