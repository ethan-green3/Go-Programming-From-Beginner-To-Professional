package main

import "fmt"

// Goal is to get from s1 declaration to Sun - Sat using just append and slice ranges

func main() {
	s1 := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	s2 := s1[len(s1)-1:]
	s2 = append(s2, s1...)
	fmt.Println(s2[:len(s2)-1])

}
