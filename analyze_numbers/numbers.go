package main

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

// func getFreq(a []int) map[int]int {
// 	return map[int]int{}
// }

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
	if a < 0 {
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

// func getMode(a []int) []int {
// 	return []int{}
// }

// func getEven(a []int) []int {
// 	return []int{}
// }

// func getOdd(a []int) []int {
// 	return []int{}
// }

// func getRanges(a []int, size int) map[string][]int {
// 	return map[string][]int{}
// }

// func getMedian(a []int) int {
// 	return 0
// }

// func getUnique(a []int) []int {
// 	return []int{}
// }

// func getAboveAvg(a []int) []int {
// 	return []int{}
// }
