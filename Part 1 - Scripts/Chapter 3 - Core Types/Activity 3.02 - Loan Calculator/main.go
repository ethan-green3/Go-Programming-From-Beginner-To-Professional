package main

import (
	"errors"
	"fmt"
)

func checkCredit(monthlyIncome int, loanAmount int, loanTerm int, creditScore int) error {

	if creditScore < 0 || monthlyIncome < 0 || loanTerm < 0 || loanAmount < 0 {
		return errors.New("Cannot have a credit score, monthly income, loan term, or loan amount less than 0")
	}
	if loanTerm%12 != 0 {
		return errors.New("Loan Term Needs to be divisible by 12")
	}

	interestRate := 0.2

	if creditScore >= 450 {
		interestRate = 0.15
	}

	monthlyPayment := ((float64(loanAmount) * interestRate) / float64(loanTerm)) + (float64(loanAmount) / float64(loanTerm))

	totalCost := (float64(loanAmount) * interestRate)
	approved := true

	if float64(monthlyIncome) <= monthlyPayment {
		approved = false
	}

	fmt.Println("Applicant")
	fmt.Println("-----------")
	fmt.Println("Credit Score: ", creditScore)
	fmt.Println("Income: ", monthlyIncome)
	fmt.Println("Loan Amount: ", loanAmount)
	fmt.Println("Loan Term: ", loanTerm)
	fmt.Println("Monthly Payment: ", monthlyPayment)
	fmt.Println("Rate: ", interestRate*100)
	fmt.Println("Total Cost ", totalCost)
	fmt.Println("Approved: ", approved)

	return nil
}

func main() {

	// checkCredit Parameters are monthly income, loan amount, loan term, credit score
	applicant1 := checkCredit(1000, 10000, 12, 350)
	if applicant1 != nil {
		fmt.Println(applicant1)
	}
	fmt.Println("")
	applicant2 := checkCredit(1000, 1000, 24, 500)
	if applicant2 != nil {
		fmt.Println(applicant2)
	}

}
