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

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
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

	defer dd.report()

	if err := dd.validateLastName(); err != nil {
		fmt.Println(err)
	}

	if err := dd.validateRoutingNumber(); err != nil {
		fmt.Println(err)
	}
}
