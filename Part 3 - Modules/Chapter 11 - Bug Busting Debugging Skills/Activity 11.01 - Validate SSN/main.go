package main

import (
	"errors"
	"log"
	"unicode"
)

var (
	ErrInvalidSSNLength  = errors.New("Error: Invalid SSN Length")
	ErrInvalidSSNNumbers = errors.New("Error: SSN has non-numeric digits")
	ErrInvalidSSNPrefix  = errors.New("Error: SSN Has invalid prefix (Leading 3 zeros)")
	ErrInvalidDigitPlace = errors.New("Error: Starts with 9 and needs a 7 or 9 in fourth place")
)

func checkSSNLength(ssn string) error {
	if len(ssn) != 11 {
		return ErrInvalidSSNLength
	}
	return nil
}

func checkSSNPrefix(ssn string) error {
	prefix := ssn[0:2]
	if prefix == "000" {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func checkSSNNumber(ssn string) error {
	for _, v := range ssn {
		if v == '-' {
			continue
		}
		if !unicode.IsDigit(v) {
			return ErrInvalidSSNNumbers
		}
	}
	return nil
}

func checkSSNDigitPlace(ssn string) error {
	firstDigit := ssn[0]
	if firstDigit == '9' {
		if ssn[4] != '7' {
			return ErrInvalidDigitPlace
		}
		if ssn[4] != '9' {
			return ErrInvalidDigitPlace
		}
	}
	return nil
}

func validateSSN(ssn string) error {
	if err := checkSSNLength(ssn); err != nil {
		log.Print(err)
	}
	if err := checkSSNDigitPlace(ssn); err != nil {
		log.Print(err)
	}
	if err := checkSSNNumber(ssn); err != nil {
		log.Print(err)
	}
	if err := checkSSNPrefix(ssn); err != nil {
		log.Print(err)
	}
	return nil
}

func main() {
	ssns := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33- 3333", "087-65-4321", "123-45-zzzz"}
	for i := range ssns {
		log.Println("Testing ssn: ", ssns[i])
		validateSSN(ssns[i])
	}
}
