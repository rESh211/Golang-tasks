package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (u *User) GetName() (string, int) { // *User - приемник метода
	return u.name, u.GetAge()
}

func (u *User) GetAge() int {
	return u.age
}

func main() {
	var num1 int = 10
	var num2 int = num1
	num2 = 20

	fmt.Println(num1)
	fmt.Println(num2)

	var slice1 = []int{1, 2, 3}
	var slice2 = slice1
	slice2[0] = 4

	fmt.Println(slice1)
	fmt.Println(slice2)

	gg := &User{name: "Ivan", age: 25}
	gg1 := new(User)
	gg1.name = "Petr"
	gg1.age = 30
	fmt.Println(gg)
	fmt.Println(gg1)

	fmt.Println(gg1.GetName())

	product := new(Product.Product)
	product.Name = "Sergey"
	fmt.Println(product)
}
