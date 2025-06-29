package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Sum of Even Numbers up to n is: %v \n\n", answer02(1000))
}

func answer02(number int) int {
	// Calculate the sum of even numbers up to a given number n.
	sum := 0
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}