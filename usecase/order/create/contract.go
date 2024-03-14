package create

type Input struct {
	ClientName string `json:"nome_do_cliente"`
	CPF        string `json:"cpf"`
	Address    string `json:"endereco"`
	Product    []int  `json:"produtos"`
}

type Output struct {
	ClientName string `json:"nome_do_cliente"`
	CPF        string `json:"cpf"`
	Address    string `json:"endereco"`
	Product    []int  `json:"produtos"`
}
