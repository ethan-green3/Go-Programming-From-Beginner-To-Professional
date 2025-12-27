package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	day := current.Format("2")
	month := current.Format("1")
	year := current.Format("2006")
	hour := current.Format("3")
	minute := current.Format("4")
	second := current.Format("5")

	fmt.Printf("%s:%s:%s %s/%s/%s\n", hour, minute, second, year, month, day)
}
