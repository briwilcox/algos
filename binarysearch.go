package main

import (
	"fmt"
)

func binSearch(searchspace []int, key int) int {
	var min, max int 
	min = searchspace[0] 
	max = searchspace[len(searchspace)-1] 
	for {
		if max < min {
			return -1
		}
		m:= (min + max)/2 
		if (searchspace[m] < key) {
			min = m + 1
		} else if (searchspace[m] > key) {
			max = m - 1
		} else {
			return m
		}
	}
}

func searchTest(a []int, key int){
	fmt.Println("Binary Search:") 
	index := binSearch(a, key) 
	if index == -1 {
		fmt.Println("Search Space: ", a)
		fmt.Println(key,"was not found")
	} else {
		fmt.Println("Search Space: ", a)
		fmt.Println("a[",index, "] = ", a[index])
	}
	fmt.Println("")
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} 
	searchTest(a, 9)
	searchTest(a, 2)
	searchTest(a, 15)
	searchTest(a, 5)
	searchTest(a, 10)
}
