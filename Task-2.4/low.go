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
