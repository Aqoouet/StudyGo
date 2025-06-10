package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type leetCodeFuncTestSuite struct {
	suite.Suite
}

func TestLeetCodeFuncSuite(t *testing.T) {
	suite.Run(t, new(leetCodeFuncTestSuite))
}

func (suite *leetCodeFuncTestSuite) TestLeetCodeFunc() {
	tests := []struct {
		name     string
		input1   []int
		input2   int
		expected []int
	}{
		{
			name:     "первый пример",
			input1:   []int{1, 2, 3, 4, 5, 6, 7},
			input2:   3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "второй пример",
			input1:   []int{-1, -100, 3, 99},
			input2:   2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "2 числа и k больше 2",
			input1:   []int{1, 2},
			input2:   3,
			expected: []int{2, 1},
		},
		{
			name:     "2 числа и k равно 2",
			input1:   []int{1, 2},
			input2:   2,
			expected: []int{1, 2},
		},
		{
			name:     "3 числа и k равно 4",
			input1:   []int{1, 2, 3},
			input2:   4,
			expected: []int{3, 1, 2},
		},
		// {
		// 	name:     "вначале нули",
		// 	input:    []int{0, 0, 0, 0, 3},
		// 	expected: 0,
		// },
		// {
		// 	name:     "отрицательные присутствуют",
		// 	input:    []int{-3, -1, -1, 0, 0, 0, 0, 0, 2},
		// 	expected: 0,
		// },
		// {
		// 	name:     "много отрицательных",
		// 	input:    []int{-3, -3, -2, -1, -1, 0, 0},
		// 	expected: -3,
		// },
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			// Копируем входной массив, чтобы не изменять исходный в тесте
			nums := make([]int, len(tt.input1))
			copy(nums, tt.input1)

			// Вызываем тестируемую функцию
			leetCodeFunc(nums, tt.input2)

			// Проверяем длину результата
			assert.Equal(suite.T(), tt.expected, nums)

			// Проверяем содержимое среза до k
			// assert.Equal(suite.T(), tt.expected, nums[:k])
		})
	}
}
