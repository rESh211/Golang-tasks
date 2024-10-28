package products

type Product struct {
	Name  string
	price float64
}

func (p *Product) GetPrice() float64 {
	return p.price
}
