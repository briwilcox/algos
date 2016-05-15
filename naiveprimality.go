package main

import (
	"fmt"
	"math"
)

func naivePrimality(isPrime float64) bool {
	if isPrime == 1 {
		return false
	}
	if isPrime == 2 {
		return true
	}
	if int(isPrime)%2 == 0 {
		return false
	}
	for d := float64(3); d <= math.Sqrt(isPrime); d++ {
		if int(isPrime)%int(d) == 0 {
			return false
		}
	}
	return true
}

func isPrime(isPrime float64){
    if naivePrimality(isPrime) {
		fmt.Println(isPrime, "is prime")
    } else {
		fmt.Println(isPrime, "is not prime")
    }
}

func main() {
	for i := 0; i <= 25; i++ {
        isPrime(float64(i))
    }
}
