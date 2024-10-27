package main

import "fmt"

type user1 struct {
	id   string
	name string
}

func Push1(s []*user1) []*user1 { // первый,добавление в начало массива
	g1 := &user1{id: "433434", name: "Pablo"}
	s = append([]*user1{g1}, s...)
	return s
}

func Append1(s []*user1) []*user1 { // второй, добавление в конец массива
	g1 := &user1{id: "747444", name: "Piter"}
	return append(s, g1)
}

func Pop1(s []*user1) []*user1 { // третий, удаление последний
	return s[:len(s)-1]
}

func Delete1(s []*user1, y string) []*user1 { // четвертый, удаление нужной нам строки
	result := make([]*user1, 0)
	for _, i := range s {
		if i.name != y {
			result = append(result, i)
		} else {
			continue
		}
	}
	return result
}

func Concat1(s []*user1) []*user1 { // пятый, объединение массивов
	g1 := &user1{id: "841532", name: "Hurumi"}
	return append(s, g1)
}

func Indexof1[t any](s []t)

func main() {
	arr := []*user1{
		&user1{id: "123534", name: "Ivan"},
		&user1{id: "432432", name: "Petr"},
		&user1{id: "564354", name: "Sergey"},
		&user1{id: "543566", name: "Alexandr"},
	}

	v := Concat1(arr)
	for _, i := range v {
		fmt.Println(i)
	}
}
