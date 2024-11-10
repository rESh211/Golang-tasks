package main

import "fmt"

type User3 struct {
	id    string
	name  string
	email string
}

func Pus3[t *User3](arr []t) []t {
	gg := &User3{id: "323233", name: "Lera", email: "y4g4@mail.ru"}
	return append([]t{gg}, arr...)
}

func Append3[t *User3](arr []t) []t {
	gg := &User3{id: "223121", name: "Micha", email: "bt321@mail.ru"}
	return append(arr, []t{gg}...)
}

func main() {
	arr3 := []*User3{
		&User3{id: "123534", name: "Ivan", email: "rr@mail.ru"},
		&User3{id: "432432", name: "Petr", email: "trvw@mail.ru"},
		&User3{id: "564354", name: "Sergey", email: "hhnwv@mail.ru"},
		&User3{id: "543566", name: "Alexandr", email: "2143@mail.ru"},
	}

	arr4 := []*User3{
		&User3{id: "123534", name: "Ivan", email: "rr@mail.ru"},
		&User3{id: "432432", name: "Petr", email: "trvw@mail.ru"},
	}

	gg1 := Pus3(arr3)
	gg2 := Append3(gg1)

	for _, i := range gg2 {
		fmt.Println(i)
	}

}
