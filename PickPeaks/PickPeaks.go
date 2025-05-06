package main

import (
	"fmt"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(arr []int) PosPeaks {

	StartPlateauInd := []int{}
	EndPlateauInd := []int{}

	for i := 0; i <= len(arr)-1; i++ {
		if i == 0 && arr[i] != arr[i+1] { // является ли первый элмент плато из 1ого элемента
			StartPlateauInd = append(StartPlateauInd, i)
			EndPlateauInd = append(EndPlateauInd, i)
		} else if i == len(arr)-1 && arr[i] != arr[i-1] {
			StartPlateauInd = append(StartPlateauInd, i)
			EndPlateauInd = append(EndPlateauInd, i)
		} else if arr[i] == arr[i+1] { // определяем длину плато для случая когда оно состоит из 2х и более элементов
			StartPlateauInd = append(StartPlateauInd, i)
			initialPos := i
			currentPos := i
			for j := i; j <= len(arr)-1 && arr[j] == arr[i]; j++ {
				currentPos++
			}
			EndPlateauInd = append(EndPlateauInd, currentPos-1)
			i += currentPos - 1 - initialPos
		} else if arr[i] != arr[i+1] && arr[i] != arr[i-1] { // добавляем плато из 1 элемента
			StartPlateauInd = append(StartPlateauInd, i)
			EndPlateauInd = append(EndPlateauInd, i)
		}

	}

	PosResult := []int{}
	PeaksResult := []int{}

	for i := 0; i <= len(StartPlateauInd)-1; i++ {
		start := StartPlateauInd[i]
		end := EndPlateauInd[i]
		if start != 0 && end != len(arr)-1 && arr[start-1] <= arr[start] && arr[end+1] <= arr[end] {
			PosResult = append(PosResult, start)
			PeaksResult = append(PeaksResult, arr[start])
		}
	}

	return PosPeaks{
		Pos:   PosResult,
		Peaks: PeaksResult,
	}
}

func main() {
	arr := []int{1, 2, 2}
	fmt.Println(PickPeaks(arr))
}
