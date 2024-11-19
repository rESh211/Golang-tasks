package task1

import "fmt"

type List[T any] struct {
	// Значение текущего элемента.
	value *T

	// Указатель на следуюзий элемент.
	next *List[T]

	// Указатель на предыдущий элемент элемент.
	prev *List[T]

	// Размер списка.
	Length int
}

func CreateList[T any]() *List[T] {
	return &List[T]{
		Length: 0,
	}
}
func (list *List[T]) ForEach(f func(val *T)) {
	current := list
	for current != nil {
		if current.value != nil { // Проверяем, что значение не nil
			f(current.value)
		}
		current = current.next
	}
}

func (list *List[T]) Filter(f func(val *T) bool) *List[T] { //Фильтрация списка и получение нового массива
	fgg := CreateList[T]()

	list.ForEach(func(value *T) {
		if f(value) {
			fgg.Push(value)
		}
	})
	return fgg
}
func (list *List[T]) IndexOf(f func(val *T) bool) int { //Получение индекса элемента по заданному условию
	c := list
	index := 0
	for c != nil { // Проходим по всем элементам списка
		if c.value != nil && f(c.value) {
			return index
		}
		c = c.next
		index++
	}
	return -1
}

func (list *List[T]) Some(f func(val *T) bool) bool { // Проверить, есть элемент по заданным условиям
	g := list.IndexOf(f)
	return g != -1
}

func (list *List[T]) Push(value *T) {
	current := list

	fmt.Println(value)

	if current.value == nil { // проверка на пустой список
		fmt.Println("Первая ячейка пустая")
		current.value = value
	} else {
		fmt.Println(current.value)
		for current.next != nil {
			fmt.Println(current.value)
			current = current.next
		}

		current.next = &List[T]{
			value: value,
			prev:  current,
		}
	}

	list.Length++

}
