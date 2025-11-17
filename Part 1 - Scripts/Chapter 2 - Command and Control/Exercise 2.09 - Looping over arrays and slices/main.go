package main

import "fmt"

func main() {
	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Modern way to use range to loop over a collection, python like syntax
	for i := range len(names) {
		fmt.Println(names[i])
	}
}
