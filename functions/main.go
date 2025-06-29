package main

import (
	"fmt"
)

func main() {
	fmt.Println(FACTORIAL(5));
	fmt.Println(ADD_TWO_NUMBERS(5 , 2));

	greet := func (username string) {
		fmt.Println("Hello " + username);
	}

	greet("Shubh");

}

func FACTORIAL(number int) int { // focus on the return type!!
	fact := 1;
	for i := 1; i <= number; i++ {
		fact *= i
	}
	return fact
}

func ADD_TWO_NUMBERS(num1, num2 int) int {
	return num1 + num2;
}