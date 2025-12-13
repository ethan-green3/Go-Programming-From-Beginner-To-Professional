package main

import (
	"fmt"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *developer) LogHours(wd Weekday, hours int) {
	d.WorkWeek[wd] = hours
}

func (d *developer) HoursWorked() int {
	hours := 0
	for _, hrs := range d.WorkWeek {
		hours += hrs
	}
	return hours
}

func nonLoggedHours() func(int) int {
	loggedSoFar := 0
	return func(h int) int {
		loggedSoFar += h
		fmt.Println("Tracking hours worked so far today: ", loggedSoFar)
		return loggedSoFar
	}
}

func (d *developer) PayDay() (int, bool) {
	hoursWorked := d.HoursWorked()
	overtimeHours := 0
	if hoursWorked > 40 {
		regularPay := 40 * d.HourlyRate
		overtimeHours = hoursWorked - 40
		overtimePay := (d.HourlyRate*2)*overtimeHours + regularPay
		return overtimePay, true
	}

	return (hoursWorked * d.HourlyRate) + overtimeHours, false
}

func (d *developer) PayDetails() {
	for i, hours := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday Hours: ", hours)
		case 1:
			fmt.Println("Monday Hours: ", hours)
		case 2:
			fmt.Println("Tuesday Hours: ", hours)
		case 3:
			fmt.Println("Wednesday Hours: ", hours)
		case 4:
			fmt.Println("Thursday Hours: ", hours)
		case 5:
			fmt.Println("Friday Hours: ", hours)
		case 6:
			fmt.Println("Saturday Hours: ", hours)
		}
	}

	fmt.Println()
	fmt.Printf("Hours worked this week: %v\n", d.HoursWorked())
	amount, overTimepay := d.PayDay()
	fmt.Printf("Pay for the week: $%d\n", amount)
	fmt.Printf("Is this overtime pay: %v\n", overTimepay)
}

func main() {
	dev := developer{}
	nonLogHours := nonLoggedHours()
	nonLogHours(2)
	nonLogHours(3)
	nonLogHours(5)

	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	dev.LogHours(Wednesday, 10)
	dev.LogHours(Thursday, 10)
	dev.LogHours(Friday, 6)
	dev.LogHours(Saturday, 8)
	dev.HourlyRate = 10

	dev.PayDetails()
}
