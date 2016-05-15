package main

import (
	"fmt"
)

func bubbleSort(toSort []int) {
	n := len(toSort)
	for a := 0; a < (n - 1); a++ {
		for b := 0; b < (n - a - 1); b++ {
			if toSort[b] > toSort[b+1] {
				temp := toSort[b]
				toSort[b] = toSort[b+1]
				toSort[b+1] = temp
			}
		}
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted: ")
	fmt.Println(a)
	bubbleSort(a)
	fmt.Println("Sorted via Bubble Sort: ")
	fmt.Println(a)

	b := []int{0, 9, 3, 5, 4, 1, 6, 7, 8, 2}
	fmt.Println("Unsorted: ")
	fmt.Println(b)
	bubbleSort(b)
	fmt.Println("Sorted via Bubble Sort: ")
	fmt.Println(b)
}
