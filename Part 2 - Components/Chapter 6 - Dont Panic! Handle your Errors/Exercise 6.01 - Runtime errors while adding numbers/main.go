package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}
	total := 0
	// index out of bounds runtime error
	for i := 0; i <= 10; i++ {
		total += nums[i]
	}
	fmt.Println("Total: ", total)
}
