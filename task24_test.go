package main

import (
	"testing"

	task24 "github.com/rESh211/golang-tasks/Task-2.4"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() {
		arr := []int{3, 5, 2, 6, 2}
		reversed := []int{2, 6, 2, 5, 3}

		arrReversed := task24.Reverse(arr)

		for index, _ := range reversed {
			assert.Equal(reversed[index], arrReversed[index], "Сравниваем элементы")
		}
	}, "Тест дозвона до функции")
}
