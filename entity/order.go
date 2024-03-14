package entity

type Order struct {
	ID         int
	ClientName string
	CPF        string
	Address    string
	Product    []int
}
