package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	day := strconv.Itoa(now.Day())
	month := now.Format("01")
	year := strconv.Itoa(now.Year())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	seconds := strconv.Itoa(now.Second())

	fmt.Printf("%s:%s:%s %s/%s/%s\n", hour, minute, seconds, year, month, day)
}
