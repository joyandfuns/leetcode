package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 8, 3, 1, 0, 22, 53, 21}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// index from 0 to len - 1, the last value needn't sort
	for i := 0; i < len(arr)-1; i++ {
		// that first element as min value
		minIndex := i
		// index from current + 1 (skip min value) to len - 1, find min index
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// swap value in current and minIndex
		swap(arr, minIndex, i)
	}
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
