package create

type Input struct {
	Nome       string  `json:"nome"`
	Tamanho    int     `json:"tamanho"`
	Cor        string  `json:"cor"`
	Quantidade int     `json:"quantidade"`
	Código     int     `json:"codigo"`
	Valor      float64 `json:"valor"`
}

type Output struct {
	ID         string  `json:"id"`
	Nome       string  `json:"nome"`
	Tamanho    int     `json:"tamanho"`
	Cor        string  `json:"cor"`
	Quantidade int     `json:"quantidade"`
	Código     int     `json:"codigo"`
	Valor      float64 `json:"valor"`
}
