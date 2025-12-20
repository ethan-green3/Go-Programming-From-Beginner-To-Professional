package main

import (
	"fmt"

	"activity10.01/pkg/payroll"
)

func init() {
	fmt.Println("Welcome the to the Employee Pay and Performance Review")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	d := payroll.Developer{Individual: payroll.Employee{Id: "123", FirstName: "Ethan", LastName: "Green"}, HourlyRate: 10, HoursWorkedInYear: 2080, Review: make(map[string]interface{})}
	d.Review["WorkQuality"] = 5
	d.Review["Teamwork"] = 2
	d.Review["Communication"] = "Poor"
	d.Review["Problem-solving"] = 4
	d.Review["Dependability"] = "Unsatisfactory"
	payroll.PayDetails(&d)
	d.DevReview()
}

func main() {

}
