package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	wordWithHighestCount := ""
	highestCount := 0

	for word, count := range words {
		if count > highestCount {
			wordWithHighestCount = word
			highestCount = count
		}
	}

	fmt.Println("Most popular word:", wordWithHighestCount)
	fmt.Println("With a count of:", highestCount)
}
