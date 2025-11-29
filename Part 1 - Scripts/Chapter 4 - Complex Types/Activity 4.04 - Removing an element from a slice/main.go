package main

import "fmt"

func main() {
	words := []string{"Good", "Good", "Bad", "Good", "Good"}
	var goodWords []string
	for _, word := range words {
		if word != "Bad" {
			goodWords = append(goodWords, word)
		}
	}
	fmt.Println(goodWords)
}
