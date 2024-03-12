package find_by_name

type Input struct {
	Nome string
}

type Output struct {
	ID       int     `json:"id"`
	Name     string  `json:"nome"`
	Size     int     `json:"tamanho"`
	Color    string  `json:"cor"`
	Quantity int     `json:"quantidade"`
	Code     int     `json:"codigo"`
	Value    float64 `json:"valor"`
}
