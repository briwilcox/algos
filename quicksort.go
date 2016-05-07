package main

import (
	"fmt"
)

// Divide portion of the routine
func partition(myList []int, start int, end int) int {
	pivot := myList[start]
	left := start + 1
	right := end
	var done bool = false
	for done == false {
		for left <= right && myList[left] <= pivot {
			left = left + 1
		}
		for myList[right] >= pivot && right >= left {
			right = right - 1
		}
		if right < left {
			done = true
		} else {
			//Swap vals
			temp := myList[left]
			myList[left] = myList[right]
			myList[right] = temp
		}
	}
	//swap start with myList[right]
	temp := myList[start]
	myList[start] = myList[right]
	myList[right] = temp
	return right
}

// Conquer portion of the routine
func quicksort(myList []int, start int, end int) []int {
	if start < end {
		//Partition the list
		pivot := partition(myList, start, end)
		//Sort both halves
		quicksort(myList, start, pivot-1)
		quicksort(myList, pivot+1, end)
	}
	return myList
}

// Pass various lists of unsorted integers to quicksort, print before and after
func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted: ")
	fmt.Println(a)
	quicksort(a, 0, len(a)-1)
	fmt.Println("Sorted via Quicksort: ")
	fmt.Println(a)

	b := []int{0, 9, 3, 5, 4, 1, 6, 7, 8, 2}
	fmt.Println("Unsorted: ")
	fmt.Println(b)
	quicksort(b, 0, len(b)-1)
	fmt.Println("Sorted via Quicksort: ")
	fmt.Println(b)
}
