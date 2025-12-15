package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (d *directDeposit) validateLastName() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if d.lastName == "" {
		panic(ErrInvalidLastName)
	}

}

func (d *directDeposit) report() {
	fmt.Printf("Last name: %v\nFirst name: %v\nBank name: %v\nRouting Number: %d\nAccount Number: %d\n", d.lastName, d.firstName, d.bankName, d.routingNumber, d.accountNumber)
}

func main() {
	dd := directDeposit{
		lastName:      "",
		firstName:     "Ethan",
		bankName:      "Test",
		routingNumber: 50,
		accountNumber: 20,
	}
	dd.validateLastName()
	dd.validateRoutingNumber()
	dd.report()

}
