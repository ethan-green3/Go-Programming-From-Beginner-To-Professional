package main

import "fmt"

func main() {

	monthlyIncome := 1000
	loanAmount := 10000
	loanTerm := 12
	goodCreditScore := false
	creditScore := 350
	interestRate := 0.2

	if creditScore < 0 || monthlyIncome < 0 || loanTerm < 0 || loanAmount < 0 {
		fmt.Println("Cannot have a credit score, monthly income, loan term, or loan amount less than 0")
		return
	}
	if loanTerm%12 != 0 {
		fmt.Println("Loan term must be divisible by 12")
		return
	}
	if creditScore >= 450 {
		goodCreditScore = true
		interestRate = 0.15
	}

	monthlyPayment := ((float64(loanAmount) * interestRate) / float64(loanTerm)) + (float64(loanAmount) / float64(loanTerm))

	if monthlyPayment > (float64(monthlyIncome)*0.2) && goodCreditScore == true {
		monthlyPayment = float64(monthlyIncome) * 0.2
	}

	if monthlyPayment > (float64(monthlyIncome)*0.1) && goodCreditScore == false {
		monthlyPayment = float64(monthlyIncome) * 0.1
	}

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

}
