package task1

type User struct {
	id    string
	name  string
	email string
	age   int
}

func Push(arr []*User) []*User { // дабавление в начала массива
	g := &User{id: "684848", name: "Ilya", email: "Ilya_Hte@gmail.com", age: 32}
	return append([]*User{g}, arr...)
}

func Append(arr []*User) []*User { // дабавление в конец массива
	g := &User{id: "357586", name: "Katya", email: "Katya81_1@gmail.com", age: 27}
	return append(arr, []*User{g}...)
}

func Pop(arr []*User) []*User { // удаление последней строчки из массива
	if len(arr) == 0 {
		return arr
	}
	return arr[:len(arr)-1]
}

func Delete(arr []*User, s string) []*User { // удаление нужного нам массива
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

func Concat(arr []*User) []*User { // объединение массивов
	g := &User{id: "864728", name: "Galya", email: "22Galya@gmail.com", age: 36}
	return append(arr, []*User{g}...)
}

/*func run(arr []*User) (errMsg string) {
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
}*/

func IndexOf[T any](arr []T, f func(val T) bool) int { //Получение индекса элемента по заданному условию
	for i, y := range arr {
		if f(y) {
			return i
		}
	}
	return -1
}

func Some[T any](arr []T, f func(val T) bool) bool { //Проверить, есть элемент по заданным условиям
	g := IndexOf(arr, f)
	return g != -1
}

func Filter[T any](arr []T, f func(val T) bool) []T { //Фильтрация списка и получение нового массива
	g1 := make([]T, 0)

	for _, i := range arr {

		if f(i) {
			g1 = append(g1, i)
		}

	}
	return g1
}

/*func Find[T any](arr []T, f func(val T) bool) []T { //Поиск элемента по условию, Filter делает тоже самое
	g := make([]T, 0)
	for _, i := range arr {
		if f(i) {
			g = append(g, i)
		}
	}
	return g
}*/

func Sort[T any](arr []T, f func(first T, second T) int) []T { //Сортировка массива и получение упорядоченного массива по заданному условию
	for i, y := range arr {
		for i1, y1 := range arr {
			hh := f(y, y1)
			if hh < 0 { // методом пузырька
				buf := arr[i]
				arr[i] = arr[i1]
				arr[i1] = buf
			}
		}
	}
	return arr
}

func ForEach[T any](arr []T, f func(val T)) { //Перебор каждого элемента массива и выполнение какой либо функции на каждом шаге
	for _, y := range arr {
		f(y)
	}
}

func Map[TIn any, TOut any](arr []TIn, f func(val TIn) TOut) []TOut { //Формирование нового массива по функции трансформации
	gg := make([]TOut, len(arr))
	for _, y := range arr {
		gg = append(gg, f(y))
	}
	return gg
}

/*func main() {// Закемментил, потому что перенес все выводы и структуру в run.go
arr := []*User{
	&User{id: "123534", name: "Ivan", email: "212Ivan@gmail.com", age: 12},
	&User{id: "432432", name: "Petr", email: "Petr.21Gd.2@mail.ru", age: 21},
	&User{id: "564354", name: "Sergey", email: "Sergey_ggF@gmail.com", age: 43},
	&User{id: "543566", name: "Alexander", email: "Alexander2@mail.ru", age: 16},
}
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
