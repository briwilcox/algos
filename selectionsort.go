package main

import (
	"fmt"
)

func selectionSort(toSort []int) {
	n := len(toSort)
	for j := 0; j < n-1; j++ {
		iMin := j
		for i := j + 1; i < n; i++ {
			if toSort[i] < toSort[iMin] {
				iMin = i
			}
		}
		if iMin != j {
			temp := toSort[j]
			toSort[j] = toSort[iMin]
			toSort[iMin] = temp
		}
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted: ")
	fmt.Println(a)
	selectionSort(a)
	fmt.Println("Sorted via Selection Sort: ")
	fmt.Println(a)

	b := []int{0, 9, 3, 5, 4, 1, 6, 7, 8, 2}
	fmt.Println("Unsorted: ")
	fmt.Println(b)
	selectionSort(b)
	fmt.Println("Sorted via Selection Sort: ")
	fmt.Println(b)
}
