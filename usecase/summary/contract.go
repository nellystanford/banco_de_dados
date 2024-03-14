package summary

type Output struct {
	TotalOrders             int     `json:"quantidade_total_pedidos"`
	TotalClients            int     `json:"clientes_registrados"`
	NewsletterSubscriptions int     `json:"inscricoes_newsletters"`
	SalesAmount             float64 `json:"valor_vendas"`
}
