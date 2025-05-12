package main

import (
	"testing"
)

var tests = []struct {
	name     string
	input    []int
	size     int
	freq     map[int]int
	min      int
	max      int
	avg      int
	mode     []int
	even     []int //четные
	odd      []int //нечетные
	ranges   map[string][]int
	median   int
	unique   []int
	aboveAvg []int
}{
	{
		name:     "стандартный вариант 1", //0
		input:    []int{1, 2, 1, 3, 2},
		size:     2,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "много больших цифр", //1
		input:    []int{10, 20, 10, 30, 40, 50, 40, 30, 20, 10},
		size:     20,
		freq:     map[int]int{10: 3, 20: 2, 30: 2, 40: 2, 50: 1},
		min:      10,
		max:      50,
		avg:      25,
		mode:     []int{10},
		even:     []int{10, 20, 10, 30, 40, 50, 40, 30, 20, 10},
		odd:      []int{},
		ranges:   map[string][]int{"10-30": {10, 20, 10, 20}, "31-50": {30, 40, 50, 40, 30}},
		median:   25,
		unique:   []int{10, 20, 30, 40, 50},
		aboveAvg: []int{30, 40, 50, 40, 30},
	},
	{
		name:     "отрицательные и положительные числа", //2
		input:    []int{-15, -3, 7, 14, 20},
		size:     5,
		freq:     map[int]int{-15: 1, -3: 1, 7: 1, 14: 1, 20: 1},
		min:      -15,
		max:      20,
		avg:      -1,
		mode:     []int{-15, -3, 7, 14, 20},
		even:     []int{-15, -3, 7},
		odd:      []int{14, 20},
		ranges:   map[string][]int{"-15--11": {-15}, "-5--1": {-3}, "0-4": {}, "5-9": {7}, "10-14": {14}},
		median:   7,
		unique:   []int{-15, -3, 7, 14, 20},
		aboveAvg: []int{7, 14, 20},
	},
	{
		name:     "все одинаковые числа", //3
		input:    []int{5, 5, 5, 5},
		size:     10,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "пустой массив", //4
		input:    []int{},
		size:     0,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "возрастающие числа", //5
		input:    []int{1, 5, 12, 15, 25},
		size:     10,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "числа по порядку", //6
		input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		size:     5,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "одно число", //7
		input:    []int{10},
		size:     10,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "отрицательные числа", //8
		input:    []int{-10, -20, -10, -30},
		size:     10,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "два числа", //9
		input:    []int{1, 3},
		size:     2,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "три числа, отрицательное, ноль и положительное", //10
		input:    []int{-5, 0, 5},
		size:     5,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "все нули", //11
		input:    []int{0, 0, 0, 0},
		size:     5,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "стандартный вариант 2", //12
		input:    []int{3, 1, 2, 1, 2, 3},
		size:     1,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "стандартный вариант 3", //13
		input:    []int{1, 1, 1, 2, 2, 3},
		size:     1,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
	{
		name:     "отрицательные и ноль", //14
		input:    []int{0, -1, -2, -3},
		size:     1,
		freq:     map[int]int{1: 2, 2: 2, 3: 1},
		min:      1,
		max:      3,
		avg:      1,
		mode:     []int{1, 2},
		even:     []int{2, 2},
		odd:      []int{1, 1, 3},
		ranges:   map[string][]int{"0-2": {1, 1}, "3-4": {2, 2, 3}},
		median:   2,
		unique:   []int{1, 2, 3},
		aboveAvg: []int{2, 2, 3},
	},
}

func TestFrequency(t *testing.T) {

}
