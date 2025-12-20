package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	quote := quote.Go()

	fmt.Println("Id: ", id)
	fmt.Println("Quote: ", quote)
}
