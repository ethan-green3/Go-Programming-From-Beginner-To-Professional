package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type server struct{}
type messageData struct {
	Message string `json:"message"`
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	messageData := messageData{}
	err := jsonDecoder.Decode(&messageData)
	if err != nil {
		log.Fatal("Error hapened in json decoder: ", err)
	}

	jsonBytes, err := json.Marshal(messageData)
	if err != nil {
		log.Fatal("Error happened in Json marshal: ", err)
	}

	log.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func main() {
	log.Fatal(http.ListenAndServe(":42069", server{}))
}
