package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type names struct {
	Names []string `json:"names"`
}

func ParseNames() (int, int) {
	r, err := http.Get("http://localhost:8085")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	names := names{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}

	electricCount := 0
	boogalooCount := 0

	for i := range names.Names {
		switch names.Names[i] {
		case "Electric":
			electricCount++
		case "Boogaloo":
			boogalooCount++
		}
	}

	return electricCount, boogalooCount
}

func main() {
	eCount, bCount := ParseNames()
	fmt.Println("Electric Count: ", eCount)
	fmt.Println("Boogaloo Count: ", bCount)
}
