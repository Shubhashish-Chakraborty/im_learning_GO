package main

import (
	"fmt"
)

func main() {
	// basic Calculator written in Go!!
	var firstnumber int
	var operation string
	var secondNumber int
	var askExit string

	for {
		fmt.Print("\n")
		fmt.Print("Enter First Number: ")
		fmt.Scan(&firstnumber)

		fmt.Print("Operation(+ , - , / , *):  ")
		fmt.Scan(&operation)

		fmt.Print("Enter Second Number: ")
		fmt.Scan(&secondNumber)

		switch operation {
		case "+":
			fmt.Printf("%v %v %v = %v \n\n", firstnumber, operation, secondNumber, firstnumber+secondNumber)
		case "-":
			fmt.Printf("%v %v %v = %v \n\n", firstnumber, operation, secondNumber, firstnumber-secondNumber)
		case "*":
			fmt.Printf("%v %v %v = %v \n\n", firstnumber, operation, secondNumber, firstnumber*secondNumber)
		case "/":
			fmt.Printf("%v %v %v = %v \n\n", firstnumber, operation, secondNumber, firstnumber/secondNumber)
		default:
			fmt.Printf("Something Went Wrong, Please Try Again! \n\n")
		}

		fmt.Print("Exit??(y/n):")
		fmt.Scan(&askExit)

		if askExit == "y" {
			break
		}
	}
}
