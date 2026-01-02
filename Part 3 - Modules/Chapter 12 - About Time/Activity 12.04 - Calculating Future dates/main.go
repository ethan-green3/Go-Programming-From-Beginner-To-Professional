package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The current time ", start.Format(time.ANSIC))
	waitTime := ((3600 * 6) + (60 * 6) + 6) * time.Second
	future := start.Add(waitTime)
	fmt.Println("6 hours, 6 minutes, and 6 seconds from now the time wil be : ", future.Format(time.ANSIC))
}
