package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("Skip")
		} else if r%2 == 0 {
			fmt.Println("STOP")
			goto STOP
		} else {
			fmt.Println(r)
		}
	}
STOP:
	fmt.Println("Goto label reached")
}
