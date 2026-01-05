package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type server struct{}
type Names struct {
	Names []string `json:"names"`
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	names := Names{}
	for {
		if len(names.Names) == 10 {
			break
		}
		num := rand.Intn(11)
		if num <= 5 {
			names.Names = append(names.Names, "Electric")
		} else {
			names.Names = append(names.Names, "Boogaloo")
		}
	}

	jsonArray, err := json.Marshal(names)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(jsonArray))
	w.Write(jsonArray)
}

func main() {
	log.Fatal(http.ListenAndServe(":8085", server{}))
}
