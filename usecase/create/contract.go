package create

type Input struct {
	Nome       string  `json:"nome"`
	Tamanho    int     `json:"tamanho"`
	Cor        string  `json:"cor"`
	Quantidade int     `json:"quantidade"`
	Codigo     int     `json:"codigo"`
	Valor      float64 `json:"valor"`
}

type Output struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tamanho    int     `json:"tamanho"`
	Cor        string  `json:"cor"`
	Quantidade int     `json:"quantidade"`
	Codigo     int     `json:"codigo"`
	Valor      float64 `json:"valor"`
}
