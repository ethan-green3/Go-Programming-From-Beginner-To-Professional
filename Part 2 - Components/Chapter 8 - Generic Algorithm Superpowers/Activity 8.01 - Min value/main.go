package main

import "fmt"

func findMinGeneric[Num int | float64](nums []Num) Num {
	var minValue Num = 999999
	for _, v := range nums {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

func main() {
	numsInt := []int{1, 32, 5, 8, 10, 11}
	numsFloat := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}
	minInt := findMinGeneric(numsInt)
	minFloat := findMinGeneric(numsFloat)
	fmt.Println("min value: ", minInt)
	fmt.Println("min value: ", minFloat)
}
