package task1

func Run() {
	list := CreateList[User]()

	list.Push(&User{id: "123534", name: "Ivan", email: "212Ivan@gmail.com", age: 12})
	list.Push(&User{id: "432432", name: "Petr", email: "Petr.21Gd.2@mail.ru", age: 21})
	list.Push(&User{id: "564354", name: "Sergey", email: "Sergey_ggF@gmail.com", age: 43})
	list.Push(&User{id: "543566", name: "Alexander", email: "Alexander2@mail.ru", age: 16})

}
