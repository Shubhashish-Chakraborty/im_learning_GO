package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Sum of Even Numbers up to n is: %v \n\n", answer02(1000))

	answer03 := func(number int) {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%v X %v = %v \n\n", number, i, number*i)
		}
	}

	answer03(4)

	fmt.Println(answer04("HELLO"))

	keepAsking := func() {
		var askNum int
		for {
			fmt.Print("Enter Number: ")
			fmt.Scan(&askNum)

			if askNum >= 0 && askNum <= 10 {
				break
			}
		}
	}

	keepAsking()
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

func answer04(theString string) string {
	// Reverse a string using a loop.
	var reversedString string = ""
	for i := len(theString) - 1; i >= 0; i-- {
		reversedString += string(theString[i])
	}
	return reversedString
}
