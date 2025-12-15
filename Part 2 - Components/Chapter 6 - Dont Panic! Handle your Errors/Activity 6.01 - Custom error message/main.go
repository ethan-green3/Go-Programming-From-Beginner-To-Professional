package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
	lastName := ""
	routingNumber := "123"
	if lastName == "" {
		fmt.Println(ErrInvalidLastName)
	}
	if len(routingNumber) < 8 {
		fmt.Println(ErrInvalidRoutingNumber)
	}
}
