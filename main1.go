package main

import (
	"fmt"
)

type User struct {
	id    string
	name  string
	email string
	age   int
}

/*func taskOne(f []*User) []*User {
	g := append(f, &User{id: "684848", name: "Ilya", email: "Ilya_Hte@gmail.com"})
	newM := User{id: "581288", name: "Dima", email: "Dima8248@mail.ru"}
	g1 := append(g, &newM)
	return g1
}*/

func run(arr []*User) (errMsg string) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg = fmt.Sprintf("Есть паника: %v", rec)
			fmt.Println(errMsg)
		}
	}()

	fmt.Println("Отправка запроса")

	index := indexOf(arr, func(user *User) bool {
		return user.email == "Sergey_ggF@gmail.com"
	})

	if index == -1 {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Printf("Элемент найден на индексе: %d\n", index)
	}

	return ""
}

func indexOf[T any](arr []T, f func(val T) bool) int {
	for i, y := range arr {
		if f(y) {
			return i
		}
	}
	return -1
}

/*func some[T any](arr []T, f func(val T) bool) bool {
	g := indexOf(arr, f)
	return g != -1
}*/

/*func Filter[T any](arr []T, f func(val T) bool) []T {
	g1 := make([]T, 0)

	for _, i := range arr {

		if f(i) {
			g1 = append(g1, i)
		}

	}
	return g1
}

func Sort[T any](arr []T, f func(first T, second T) int) []T {
	for i, y := range arr {
		for i1, y1 := range arr {
			hh := f(y, y1)
			if hh < 0 {
				buf := arr[i]
				arr[i] = arr[i1]
				arr[i1] = buf
			}
		}
	}
	return arr
}

func ForEach[T any](arr []T, f func(val T)) {
	for _, y := range arr {
		f(y)
	}
}

func Map[TIn any, TOut any](arr []TIn, f func(val TIn) TOut) []TOut {
	gg := make([]TOut, len(arr))
	for _, y := range arr {
		gg = append(gg, f(y))
	}
	return gg
}*/

func main() {
	arr := []*User{
		&User{id: "123534", name: "Ivan", email: "212Ivan@gmail.com", age: 12},
		&User{id: "432432", name: "Petr", email: "Petr.21Gd.2@mail.ru", age: 21},
		&User{id: "564354", name: "Sergey", email: "Sergey_ggF@gmail.com", age: 43},
		&User{id: "543566", name: "Alexander", email: "Alexander2@mail.ru", age: 16},
	}
	//f1 := taskOne(arr)

	/*for _, user := range f1 {
		fmt.Println(user)
	}*/
	fmt.Println(indexOf(arr, func(user *User) bool { return user.email == "Sergey_ggF@gmail.com" }))
	err := run(arr)
	if err != "" {
		fmt.Println("Ошибка:", err)
	}
	/*fmt.Println(some(arr, func(user *User) bool { return user.name == "Alexander" }))
	fmt.Println(Filter(arr, func(user *User) bool { return strings.Contains(user.email, "gmail.com") }))
	gmailuser := Filter(arr, func(user *User) bool { return strings.Contains(user.email, "gmail.com") })
	fmt.Println(gmailuser)

	for _, user := range gmailuser {
		fmt.Println(user.email)
	}
	hr := Sort(arr, func(user1 *User, user2 *User) int { return user1.age - user2.age })

	for _, user := range hr {
		fmt.Println(user.name)
	}
	ForEach(arr, func(user *User) { user.email = strings.ToLower(user.email) })

	for _, user := range arr {
		fmt.Println(user.email)
	}

	ht := Map(arr, func(user *User) string { return user.name + " " + user.email })
	fmt.Println(ht)
	*/
}
