package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Start: ", start.Format(time.ANSIC))
	time.Sleep(2 * time.Second)
	elasped := time.Since(start)
	fmt.Println("Elapsed: ", elasped)
}
