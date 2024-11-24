package task1

import "fmt"

type User struct {
	id    string
	name  string
	email string
	age   int
}

// Добавление в начало массива.
func Push(arr []*User, arr1 []*User) []*User {
	return append(arr1, arr...)
}

// Добавление в конец массива.
func Append(arr []*User, arr1 []*User) []*User {
	return append(arr, arr1...)
}

// Удаление последней строчки из массива.
func Pop(arr []*User) []*User {
	if len(arr) == 0 {
		return arr
	}
	return arr[:len(arr)-1]
}

// Удаление последней строчки из массива.
func Delete(arr []*User, s string) []*User {
	g := make([]*User, 0)
	for _, i := range arr {
		if i.id == s {
			continue
		} else {
			g = append(g, i)
		}
	}
	return g
}

// Объединение массивов.
func Concat(arr []*User, arr1 []*User) []*User {

	return append(arr, arr1...)
}

func Run1(arr []*User) (errMsg string) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg = fmt.Sprintf("Есть паника: %v", rec)
			fmt.Println(errMsg)
		}
	}()

	fmt.Println("Отправка запроса")

	index := IndexOf(arr, func(user *User) bool {
		return user.email == "Sergey_ggF@gmail.com"
	})

	if index == -1 {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Printf("Элемент найден на индексе: %d\n", index)
	}

	return ""
}

//Получение индекса элемента по заданному условию.
func IndexOf[T any](arr []T, f func(val T) bool) int {
	for i, y := range arr {
		if f(y) {
			return i
		}
	}
	return -1
}

//Проверят, есть элемент по заданным условиям.
func Some[T any](arr []T, f func(val T) bool) bool {
	g := IndexOf(arr, f)
	return g != -1
}

//Фильтрация списка и получение нового массива.
func Filter[T any](arr []T, f func(val T) bool) []T {
	g1 := make([]T, 0)

	for _, i := range arr {

		if f(i) {
			g1 = append(g1, i)
		}

	}
	return g1
}

//Поиск элемента по условию, Filter делает тоже самое.
func Find[T any](arr []T, f func(val T) bool) []T {
	g := make([]T, 0)
	for _, i := range arr {
		if f(i) {
			g = append(g, i)
		}
	}
	return g
}

//Сортировка массива и получение упорядоченного массива по заданному условию.
func Sort[T any](arr []T, f func(first T, second T) int) []T {
	for i, y := range arr {
		for i1, y1 := range arr {
			hh := f(y, y1)
			// Сортировка пузырьком.
			if hh < 0 {
				buf := arr[i]
				arr[i] = arr[i1]
				arr[i1] = buf
			}
		}
	}
	return arr
}

//Перебор каждого элемента массива и выполнение какой либо функции на каждом шаге.
func ForEach[T any](arr []T, f func(val T)) {
	for _, y := range arr {
		f(y)
	}
}

//Формирование нового массива по функции трансформации.
func Map[TIn any, TOut any](arr []TIn, f func(val TIn) TOut) []TOut {
	gg := make([]TOut, len(arr))
	for _, y := range arr {
		gg = append(gg, f(y))
	}
	return gg
}

/*func Run() {// Закемментил, потому что перенес все выводы и структуру в run.go
//f1 := taskOne(arr)

/*for _, user := range f1 {
	fmt.Println(user)
}
fmt.Println(IndexOf(arr, func(user *User) bool { return user.email == "Sergey_ggF@gmail.com" }))
err := run(arr)
if err != "" {
	fmt.Println("Ошибка:", err)
}*/
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
fmt.Println(ht)*/
