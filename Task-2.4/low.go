package task24

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
	"golang.org/x/text/number"
)

// Тип Queue с методами Add и Delete.
type Queue[T any] struct {
	element []T
	size    int
}

func (queue *Queue[T]) Add(element T) {
	queue.element = append(queue.element, element)
	queue.size++
}

func (queue *Queue[T]) Delete() (T, bool) {
	if queue.size == 0 {
		fmt.Println("Пусто")
		var zero T
		return zero, false
	}
	val := queue.element[0]
	queue.element = queue.element[1:]
	queue.size--
	return val, true
}

// Метод для получения размера очереди
func (queue *Queue[T]) Size() int {
	return queue.size
}

// Метод пуста ли очередь
func (queue *Queue[T]) Empty() bool {
	return queue.size == 0
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
func Reverse[T any](arr []T) []T {
	index := len(arr) - 1
	var reversed []T

	for i := index; i >= 0; index-- {
		reversed = append(reversed, arr[i])
	}

	return reversed
}

// Выводит элементы массива на консоль.
func PrintArray[T any](arr []T, word string) {
	for i, y := range arr {
		if i > 0 {
			fmt.Println(", ")
		}
		fmt.Println(y)
	}
	fmt.Print(word)
}

// Вычисляет сумму элементов в массиве (ограничение на числовые типы).
func Sum[T constraints.Ordered](arr []T) T {
	var numder T
	if arr != nil && len(arr) == 0 {
		return numder
	}
	for _, y := range arr {
		numder += y
	}
	return numder
}

// Тип LinkedList с методами добавления, удаления и поиска элементов.
type LinkedList[T any] struct {
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
func (linkedList *LinkedList[T]) IndexOf(arr []T, element int) int {
	for i, y := range arr {
		if element == y {
			return i
		}
	}
	return -1
}

// Тип Map, который хранит пары ключ-значение (ключ - comparable, значение - любое).
type Map[T constraints.Ordered, meaning any] struct {
	key map[T]meaning
}

//Метод добавления в Мап
func (map *Map[T, meaning])Add(one T, two meaning){
	map.key[one]=two
}

func(map *Map[T,meaning])Get(number T)(meaning,bool){
	number, text:=map.key[number]
	return number, text
}

// Напишите дженерическую функцию Sort, которая сортирует массив (ограничение на упорядоченные типы).
func Sort[T constraints.Ordered](arr []T, sorting func(a, b T) bool) []T {
	sort.Slice(arr, func(i, j int) bool {
		return sorting(arr[i], arr[j])
	})
	return arr
}
