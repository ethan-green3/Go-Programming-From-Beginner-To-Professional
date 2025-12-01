package main

import "fmt"

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

func main() {
	dev := developer{}
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday: ", dev.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", dev.WorkWeek[Tuesday])
	fmt.Println("Hours this week: ", dev.HoursWorked())
}
