package task24

import "golang.org/x/exp/constraints"

func FindMin[T constraints.Ordered](arr []T) T {
	min := arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

func Reverse[T any](arr []T) []T {
	index := len(arr) - 1
	var reversed []T

	for index >= 0 {
		reversed = append(reversed, arr[index])
		index--
	}

	return reversed
}
