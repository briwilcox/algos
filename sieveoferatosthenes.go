package main

import (
	"fmt"
	"math"
	"sort"
)

func sieveOfEratosthenes(upperBound int) {
	// Use hashmap to prevent duplicates
	primes := map[int]int{}

	upperBoundSquareRootFl := math.Sqrt(float64(upperBound))
	upperBoundSquareRoot := int(upperBoundSquareRootFl)
	isComposite := []bool{}
	for i := 0; i < upperBound; i++ {
		isComposite = append(isComposite, false)
	}
	for m := 2; m <= upperBoundSquareRoot; m++ {
		if !isComposite[m] {
			primes[m] = m
			for k := m * m; k <= int(upperBound); k += m {
				if k <= len(isComposite)-1 {
					isComposite[k] = true
				}
			}
		}
	}
	for m := upperBoundSquareRoot; m <= int(upperBound); m++ {
		if m <= len(isComposite)-1 {
			if !isComposite[m] {
				primes[m] = m
			}
		}
	}

	// Get all keys out of hashmap
	keys := []int{}
	for key := range primes {
		keys = append(keys, key)
	}

	// Sort Keys
	sort.Ints(keys)

	// Print array of all prime numbers up to upperBound
	fmt.Println(keys)
}

func main() {
	sieveOfEratosthenes(120)
}
