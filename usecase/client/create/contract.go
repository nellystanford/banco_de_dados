package create

type Input struct {
	Name       string `json:"nome"`
	Email      string `json:"email"`
	CPF        string `json:"cpf"`
	Address    string `json:"endereco"`
	Newsletter bool   `json:"newsletter"`
}

type Output struct {
	ID         int    `json:"id"`
	Name       string `json:"nome"`
	Email      string `json:"email"`
	CPF        string `json:"cpf"`
	Address    string `json:"endereco"`
	Newsletter bool   `json:"newsletter"`
}
