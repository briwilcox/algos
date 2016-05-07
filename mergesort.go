package main

import (
	"fmt"
)

func mergeSortTopDown(toSort []int, temp []int) {
	m_sort(toSort, 0, len(toSort), temp)
}

func m_sort(toSort []int, left int, right int, temp []int) {
	if right-left < 2 || len(toSort) < 2 {
		return
	}
	if left < right {
		mid := (left + right) / 2
		m_sort(toSort, left, mid, temp)
		m_sort(toSort, mid, right, temp)
		topDownMerge(toSort, left, mid, right, temp)
		copyArray(temp, left, right, toSort)
	}

}

func topDownMerge(toSort []int, left int, mid int, right int, temp []int) {
	i := left
	j := mid
	for k := left; k < right; k++ {
		if i < mid && (j >= right || toSort[i] <= toSort[j]) {
			temp[k] = toSort[i]
			i = i + 1
		} else {
			temp[k] = toSort[j]
			j = j + 1
		}
	}
}

func copyArray(temp []int, left int, right int, toSort []int) {
	for k := left; k < right; k++ {
		toSort[k] = temp[k]
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted: ")
	fmt.Println(a)
	temp := make([]int, 10)
	mergeSortTopDown(a, temp)
	fmt.Println("Sorted via MergeSort: ")
	fmt.Println(a)

	b := []int{0, 9, 3, 5, 4, 1, 6, 7, 8, 2}
	fmt.Println("Unsorted: ")
	fmt.Println(b)
	temp2 := make([]int, 10)
	mergeSortTopDown(b, temp2)
	fmt.Println("Sorted via MergeSort: ")
	fmt.Println(b)
}
