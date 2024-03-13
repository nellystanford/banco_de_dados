package entity

type Order struct {
	ID      int
	Number  int
	Client  []Client
	Address string
	Product []Product
}
