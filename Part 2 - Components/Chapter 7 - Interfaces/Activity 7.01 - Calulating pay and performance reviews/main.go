package main

import (
	"fmt"
)

type Employee struct {
	Id        string
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salarly        float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d *Developer) Pay() (string, float64) {
	return d.Individual.FirstName + " " + d.Individual.LastName, d.HourlyRate * d.HoursWorkedInYear
}

func (m *Manager) Pay() (string, float64) {
	return m.Individual.FirstName, m.Salarly + (m.Salarly * m.CommissionRate)
}

func payDetails(p Payer) {
	name, pay := p.Pay()
	fmt.Printf("Full name: %v\nPay: %.2f", name, pay)
}

func (d *Developer) devReview() float64 {
	totalScore := 0.0
	for _, v := range d.Review {
		switch val := v.(type) {
		case string:
			totalScore += float64(convertStringToInt(val))
		case int:
			totalScore += float64(val)
		default:
			fmt.Printf("Unsupported type: %T\n", v)
		}
	}

	//fmt.Printf("Total Score: %v\nLength of review map:%.2f\n", totalScore, float64(len(d.Review)))
	return totalScore / float64(len(d.Review))
}

func convertStringToInt(str string) int {
	switch str {
	case "Excellent":
		return 5
	case "Good":
		return 4
	case "Fair":
		return 3
	case "Poor":
		return 2
	case "Unsatisfactory":
		return 1
	default:
		fmt.Println("Invalid string")
		return 0
	}
}

func main() {
	d := Developer{Individual: Employee{Id: "123", FirstName: "Ethan", LastName: "Green"}, HourlyRate: 10, HoursWorkedInYear: 2080, Review: make(map[string]interface{})}
	d.Review["WorkQuality"] = 5
	d.Review["Teamwork"] = 2
	d.Review["Communication"] = "Poor"
	d.Review["Problem-solving"] = 4
	d.Review["Dependability"] = "Unsatisfactory"

	r := d.devReview()
	fmt.Printf("%v %v got a review rating of %.2f\n", d.Individual.FirstName, d.Individual.LastName, r)
	name, pay := d.Pay()
	fmt.Printf("%v got paid %.2f for the year\n", name, pay)
}
