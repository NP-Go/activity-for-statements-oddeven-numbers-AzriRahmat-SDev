package main

import "fmt"

func main() {
	//Insert code here
	var userInput1 int

	for {
		fmt.Println("Please enter a number")
		fmt.Scanln(&userInput1)

		if userInput1%2 == 0 {
			fmt.Println("This is a even number")
		} else {
			fmt.Println("This is a odd number")
		}
	}
}
