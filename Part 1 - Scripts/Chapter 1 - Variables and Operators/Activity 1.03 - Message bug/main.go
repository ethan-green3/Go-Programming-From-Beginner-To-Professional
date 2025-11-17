package main

import "fmt"

func main() {
	count := 5
	if count > 5 {
		message := "Greater than 5"
		fmt.Println(message)
	} else {
		message := "Not greater than 5"
		fmt.Println(message)
	}
	// Original snippet has print statment here and it is out of scope, print statements moved inside of if statements to be in scope
	// fmt.Println(message)
}
