package task24

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// тип Queue с методами Enqueue и Dequeue.
type Queue[T constraints.Ordered] struct {
	element []T
}

func (queue *Queue[T]) Enqueue(element T) {
	queue.element = append(queue.element, element)

}

func (queue *Queue[T]) Dequeue() T {
	if len(queue.element) == 0 {
		fmt.Println("Пусто")
	}
	val := queue.element[0]
	queue.element = queue.element[1:]
	return val
}

// Находит минимальное значение в массиве.
func FindMin[T constraints.Ordered](arr []T) T {
	min := arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

// Меняет порядок элементов в массиве.
func Reverse[T constraints.Ordered](arr []T) []T {
	index := len(arr) - 1
	var reversed []T

	for index >= 0 {
		reversed = append(reversed, arr[index])
		index--
	}

	return reversed
}

// Выводит элементы массива на консоль.
func PrintArray[T constraints.Ordered](arr []T) {
	for _, y := range arr {
		fmt.Println(y)
	}
}

// Вычисляет сумму элементов в массиве (ограничение на числовые типы).
func Sum[T constraints.Ordered](arr []T) T {
	var numder T
	for _, y := range arr {
		numder += y
	}
	return numder
}

// Тип LinkedList с методами добавления, удаления и поиска элементов.
type LinkedList[T constraints.Ordered] struct {
	//Указатель на следующий элемент.
	next *LinkedList[T]
	//Хранимое значение.
	value T
}

// Добавление элемента в массив.
func (linkedList *LinkedList[T]) Add(item T) {
	next := linkedList
	for next.next != nil {
		next = next.next
	}
	next.next = &LinkedList[T]{value: item}
}

// Удаление элемента из массива.
func (linkedList *LinkedList[T]) Get(index int) T {
	next := linkedList
	number := 0
	for next.next != nil && number <= index {
		next = next.next
		number++
	}
	return next.value
}

// Поиск элемента в массиве.
func (linkedList *LinkedList[T]) Search(element T) bool {
	next := linkedList
	if next != nil {
		if next.value == element {
			return true
		}
	}
	return false
}

// Поиск элемента в массиве.
func ElementSearch[T constraints.Ordered](arr []T, element T) bool {
	for _, y := range arr {
		if y == element {
			return true
		}
	}
	return false
}

// Тип Map, который хранит пары ключ-значение (ключ - comparable, значение - любое).
type Map[T constraints.Ordered, meaning any] struct {
	key map[T]meaning
}
