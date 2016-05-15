package main

import (
	"fmt"
)

func insertionSort(toSort []int) {
      for i := 1; i < len(toSort); i++ {
            newValue := toSort[i];
            j := i;
            for j > 0 && toSort[j - 1] > newValue {
                  toSort[j] = toSort[j - 1];
                  j--;
            }
            toSort[j] = newValue;
      }
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Unsorted: ")
	fmt.Println(a)
	insertionSort(a)
	fmt.Println("Sorted via Insertion Sort: ")
	fmt.Println(a)

	b := []int{0, 9, 3, 5, 4, 1, 6, 7, 8, 2}
	fmt.Println("Unsorted: ")
	fmt.Println(b)
	insertionSort(b)
	fmt.Println("Sorted via Insertion Sort: ")
	fmt.Println(b)
}
