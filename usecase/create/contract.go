package create

type Input struct {
	Nome       string `json:"nome,required"`
	Tamanho    int    `json:"tamanho,required"`
	Cor        string `json:"cor,required"`
	Quantidade int    `json:"quantidade,required"`
	Código     string `json:"codigo,required"`
}

type Output struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Tamanho    int    `json:"tamanho"`
	Cor        string `json:"cor"`
	Quantidade int    `json:"quantidade"`
	Código     string `json:"codigo"`
}
