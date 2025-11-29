package main

import (
	"fmt"
	"os"
)

type locale struct {
	language string
	country  string
}

var locales []locale

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need command for script, language, and country in format of _<country initials>")
		os.Exit(1)
	}

	locales = append(locales,
		locale{"en", "_US"},
		locale{"en", "_CN"},
		locale{"fr", "_CN"},
		locale{"fr", "_CN"},
		locale{"ru", "_RU"})

	input_language := os.Args[1]
	input_country := os.Args[2]

	localeToCheck := locale{input_language, input_country}

	for _, local := range locales {
		if local == localeToCheck {
			fmt.Println("Locale passed is supported")
			os.Exit(0)
		}
	}

	fmt.Println("Local passed is not supported")
	os.Exit(1)

}
