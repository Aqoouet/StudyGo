package main

import "fmt"

func leetCodeFunc(nums []int, k int) {

	if k == len(nums) {
		return
	}

	if k >= len(nums) {
		k = k % len(nums)
	}

	newInd := indexesNew(len(nums), k)

	ini := make([]int, len(nums))

	copy(ini, nums)

	for i, _ := range ini {
		nums[i] = ini[newInd[i]]
	}

}

func indexesNew(l, k int) []int {

	a := []int{}

	for i := k; i > 0; i-- {
		a = append(a, l-i)
	}

	for i := k; i < l; i++ {
		a = append(a, i-k)
	}

	return a

}

func main() {
	a := []int{1, 2, 3}
	leetCodeFunc(a, 1)
	fmt.Println(a)
}
