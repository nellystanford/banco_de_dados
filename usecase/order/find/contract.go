package find

type Output struct {
	ID         int    `json:"id"`
	ClientName string `json:"nome_do_cliente"`
	CPF        string `json:"cpf"`
	Address    string `json:"endereco"`
	Product    []int  `json:"produtos"`
}
