package main

import "fmt"

type User1 struct {
	id   string
	name string
}

func Push[t *User1](arr []t) []t { // Обобщенная функция, принимающая срез указателей на User1
	gg := &User1{id: "897456", name: "Masha"}
	return append([]t{gg}, arr...)
}
func Append[t *User1](arr []t) []t {
	gg := &User1{id: "897456", name: "Masha"}
	return append(arr, []t{gg}...)
}

func Pop[t *User1](arr []t) []t { // удаляем конец массива
	return arr[:len(arr)-1]
}

func Delete(arr []*User1, s string) []*User1 { //удаление
	gg := make([]*User1, 0)
	for _, i := range arr {
		if i.id != s {
			gg = append(gg, i)
		}
	}
	return gg

}

func Concat[t *User1](arr []t, arr1 []t) []t { // объединение
	return append(arr, arr1...)
}

func Indexof(arr []*User1, s string) []*User1 {
	gg := make([]*User1, 0)
	for _, i := range arr {
		if i.id == s {
			gg = append(gg, i)
		}
	}
	return gg
}

func Some(arr []*User1, s string) bool {
	for _, i := range arr {
		if i.id == s {
			return true
		}
	}
	return false
}

func main() {
	arr := []*User1{
		&User1{id: "123534", name: "Ivan"},
		&User1{id: "432432", name: "Petr"},
		&User1{id: "564354", name: "Sergey"},
		&User1{id: "543566", name: "Alexandr"},
	}

	/*ff := Indexof(arr, "432432")
	for _, i := range ff {
		fmt.Println(i)
	}*/

	g1 := Some(arr, "43243")
	fmt.Println(g1)

}
