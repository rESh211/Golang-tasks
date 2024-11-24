package products

import "fmt"

type Product struct {
	Name  string
	price float64
}

type Contractor struct {
	User
	ID string
}

type Employee1 struct {
	User
	ID string
}

type Authorizable interface {
	Authorize() bool
}

func (c Contractor) Authorizable() bool {
	fmt.Println("Авторизация подрядчика:", c.name)
	return c.age == 30
}
func (c Employee1) Authorizable() bool {
	fmt.Println("Авторизация пользователя:", c.name)
	return c.age == 21
}

func (p *Product) GetPrice() float64 {
	return p.price
}
