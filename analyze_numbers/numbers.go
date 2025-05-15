package main

import (
	"fmt"
	"strconv"
)

func sortSlice(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	middle := []int{a[len(a)-1]}
	p := a[len(a)-1]
	left, right := []int{}, []int{}
	arr := a[:len(a)-1]
	for _, i := range arr {
		switch {
		case i < p:
			left = append(left, i)
		case i > p:
			right = append(right, i)
		case i == p:
			middle = append(middle, i)
		}
	}
	return append(append(sortSlice(left), middle...), sortSlice(right)...)
}

func getFreq(a []int) map[int]int {

	if len(a) == 0 {
		return map[int]int{}
	}

	var rez = map[int]int{}

	for _, i := range a {
		rez[i] += 1
	}

	return rez
}

func getMin(a []int) int {

	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]
	}

	sorted := sortSlice(a)
	return sorted[0]
}

func getMax(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]
	}

	sorted := sortSlice(a)
	return sorted[len(a)-1]
}

func roundDown(a float64) int {
	if float64(int(a)) == a {
		return int(a)
	} else if a < 0 {
		return int(a) - 1
	} else if a > 0 {
		return int(a)
	} else {
		return 0
	}

}

func getAvg(a []int) int {

	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]
	}

	s := 0.0

	for _, i := range a {
		s += float64(i)
	}

	count := float64(len(a))

	return roundDown(s / count)
}

func getMode(a []int) []int {

	if len(a) == 0 {
		return []int{}
	}

	freq := getFreq(a)

	count := 0

	for _, i := range freq {
		if i > count {
			count = i
		}
	}

	var rez = []int{}

	for key, i := range freq {
		if i == count {
			rez = append(rez, key)
		}
	}

	return sortSlice(rez)
}

func getEven(a []int) []int {

	if len(a) == 0 {
		return []int{}
	}

	var rez = []int{}

	for _, i := range a {
		if i%2 == 0 {
			rez = append(rez, i)
		}
	}

	return rez
}

func getOdd(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	var rez = []int{}

	for _, i := range a {
		if i%2 != 0 {
			rez = append(rez, i)
		}
	}

	return rez
}

func getStartEnd(i int, size int) (int, int) {

	var start int

	if i >= 0 {

		start = (i / size) * size

	} else {

		start = ((i+1)/size)*size - size

	}

	end := start + size - 1

	return start, end

}

func getRanges(a []int, size int) map[string][]int {

	if len(a) == 0 || size == 0 {
		return map[string][]int{}
	}

	rez := map[string][]int{}

	for _, i := range a {

		start, end := getStartEnd(i, size)
		key1 := []rune(strconv.Itoa(start))
		key2 := []rune(strconv.Itoa(end))
		d := []rune("-")
		key := string(append(append(key1, d...), key2...))
		rez[key] = append(rez[key], i)

	}

	return rez
}

func getMedian(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]
	}

	sorted := sortSlice(a)

	var med float64

	if len(a)%2 == 0 {
		left := float64(sorted[len(a)/2-1])
		right := float64(sorted[len(a)/2])
		med = (left + right) / 2
		return roundDown(med)
	} else {
		return sorted[len(a)/2]
	}

}

func getUnique(a []int) []int {

	if len(a) == 0 {
		return []int{}
	}

	freq := getFreq(a)

	var rez = []int{}

	for key, _ := range freq {
		rez = append(rez, key)
	}

	return sortSlice(rez)
}

func getAboveAvg(a []int) []int {

	if len(a) == 0 {
		return []int{}
	}

	avg := getAvg(a)

	var rez = []int{}

	for _, i := range a {
		if i >= avg {
			rez = append(rez, i)
		}
	}

	return rez
}

func analyzeNumbers(a []int, size int) string {

	return fmt.Sprintf("Частота: %v\nМинимум: %v, Максимум: %v, Среднее: %v\nМода: %v\nЧётные: %v, Нечётные: %v\nДиапазоны (размер %v): %v\nМедиана: %v\nУникальные: %v\nЧисла выше среднего: %v\n",
		getFreq(a), getMin(a), getMax(a), getAvg(a), getMode(a), getEven(a), getOdd(a), size, getRanges(a, size), getMedian(a), getUnique(a), getAboveAvg(a))

}

func main() {
	fmt.Println(analyzeNumbers([]int{1, 2, 1, 3, 2}, 2))
}
