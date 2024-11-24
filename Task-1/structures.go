package task1

/*-import "fmt"

type CustomUser[T any] struct {
	name        string
	chief       *CustomUser[T]
	linkedUsers []T
}

type Contractor struct { // поддерживает функциональность нашего пользователя
	CustomUser[any]
	name string
}

func (user *CustomUser) GetName() string { // функция будет выводить только имя пользователя
	return user.name
}
func (user *CustomUser) Getchief() *CustomUser { // мы ссылаемся на прошлую функцию
	return user.chief
}

func Runstruct() {
	chief := &CustomUser{ // мы ссылаемся на прошлую функцию
		name: "ivan",
	}
	user := &CustomUser{
		name:  "Petr",
		chief: chief, // мы ссылаемся на прошлую функцию
	}
	user.GetName()

	fmt.Println(user.GetName())
	fmt.Println(user.Getchief().GetName()) // мы ссылаемся на прошлую функцию
}*/
