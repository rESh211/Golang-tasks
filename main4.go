package main

import "fmt"

type user2 struct {
	id   string
	name string
}

func push2[t any](s []t) []t { // первая, добавление в начало
	g := &user2{id: "828464", name: "Vasia"}
	return append([]t{any(g).(t)}, s...)
}

func append2[t any](s []t) []t { // второй тип
	g := &user2{id: "132484", name: "Rock"}
	return append(s, any(g).(t))
}

func delete2[t any](s []t) []t { // третий тип
	return s[:len(s)-1]
}

func concat2[t any](s []t) []t { // четвертый тип
	g := &user2{id: "132484", name: "Rock"}
	return append(s, any(g).(t))
}

func indexof2[t any](s []t, user2 *user2) int {
	for y, i := range s {
		if user2 == i {
			return y
		}
	}
	return -1
}

func main() {
	arr := []*user2{
		&user2{id: "123534", name: "Ivan"},
		&user2{id: "432432", name: "Petr"},
		&user2{id: "564354", name: "Sergey"},
		&user2{id: "543566", name: "Alexandr"},
	}

	fmt.Println(indexof2(arr, &user2{id: "123534"}))

	tt1 := concat2(arr)
	for _, i := range tt1 {
		fmt.Println(i)
	}
}
