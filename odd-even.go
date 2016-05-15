package main

import (
	"fmt"
)

func oddeven(toSort []int) {
	n := len(toSort)
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			for j := 2; j < n; j += 2 {
				if toSort[j] < toSort[j-1] {
					temp := toSort[j]
					toSort[j] = toSort[j-1]
					toSort[j-1] = temp
				}
			}
		} else {
			for j := 1; j < n; j += 2 {
				if toSort[j] < toSort[j-1] {
					temp := toSort[j]
					toSort[j] = toSort[j-1]
					toSort[j-1] = temp
				}
			}
		}
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted: ")
	fmt.Println(a)
	oddeven(a)
	fmt.Println("Sorted via Selection Sort: ")
	fmt.Println(a)

	b := []int{0, 9, 3, 5, 4, 1, 6, 7, 8, 2}
	fmt.Println("Unsorted: ")
	fmt.Println(b)
	oddeven(b)
	fmt.Println("Sorted via Selection Sort: ")
	fmt.Println(b)
}
