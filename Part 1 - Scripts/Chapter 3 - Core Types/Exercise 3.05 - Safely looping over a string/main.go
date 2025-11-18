package main

import "fmt"

func main() {
	logLevel := "中くらい"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
