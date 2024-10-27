func test2() {
	arr := []int{-1, -3, -5, 4, 5, 8}

	// index := indexOf(arr, isPositive)
	index := indexOf(arr, func(val int) bool {
		return val > 0
	})

	includes := some(arr, func(val int) bool {
		return val > 0
	})

	fmt.Println(index)
	fmt.Println(includes)

	positiveDigits := Filter(arr, isPositive)
	fmt.Println(positiveDigits)

	fmt.Println(Find[int](arr, isPositive))

	fmt.Println(Sort[int](arr, func(a int, b int) int {
		return a - b
	}))

}

func isPositive(val int) bool {
	return val > 0
}

func indexOf[T any](arr []T, f func(val T) bool) int {
	for i, j := range arr {
		matched := f(j)

		if matched {
			return i
		}
	}
	return -1
}

func some[T any](arr []T, f func(val T) bool) bool {
	index := indexOf(arr, f)
	return index != -1
}

func Filter[T any](arr []T, f func(val T) bool) []T {
	newArr := make([]T, 0)

	for _, j := range arr {
		matched := f(j)

		if matched {
			newArr = append(newArr, j)
		}
	}

	return newArr
}

func Find[T any](arr []T, f func(val T) bool) T {
	index := indexOf[T](arr, f)
	if index != -1 {
		return arr[index]
	} else {
		panic("No such element")
	}
}

func Sort[T any](arr []T, f func(first T, second T) int) []T {
	for indexI, valI := range arr {
		for indexJ, valJ := range arr {
			compareRes := f(valI, valJ)
			if compareRes < 0 {
				buf := arr[indexI]
				arr[indexI] = arr[indexJ]
				arr[indexJ] = buf
			}
		}
	}
	return arr
}

func ForEach[T any](arr []T, f func(val T)) {
	for _, j := range arr {
		f(j)
	}
}

func Map[TIn any, TOut any](arr []TIn, f func(val TIn) TOut) []TOut {
	newArr := make([]TOut, len(arr))
	for _, j := range arr {
		newArr = append(newArr, f(j))
	}

	return newArr
}