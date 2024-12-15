package task24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		panic("Test panic")
	})
}

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		arr := []int{3, 5, 2, 6, 2}
		reversed := []int{2, 6, 2, 5, 3}

		arrReversed := Reverse(arr)

		for index, _ := range reversed {
			assert.Equal(reversed[index], arrReversed[index], "Сравниваем элементы")
		}
	}, "Тест дозвона до функции")
}
