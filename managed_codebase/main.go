package main

import (
	"fmt";
	"myGoProject/helper";
)

func main() {
	var askAgain string;
	for {
		fmt.Print("want to generate a Random Number(y/n): ");
		fmt.Scan(&askAgain);
		
		if askAgain == "n" {
			fmt.Println("SEE YOU SOON!");
			break;
		}
		
		fmt.Printf("Here is your number: %v \n\n" , helper.RandomNumber());
	}
}