package main

import (
	"fmt"
	"os"
)

func main() {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
	if len(os.Args) < 2 {
		fmt.Println("Need to input a user's ID")
		os.Exit(1)
	}

	input := os.Args[1]
	name, ok := users[input]
	if !ok {
		fmt.Println("No user found with the passed ID")
		os.Exit(1)
	}

	fmt.Println("Id passed in:", input, "User:", name)
	fmt.Printf("Hello %v!", name)
}
