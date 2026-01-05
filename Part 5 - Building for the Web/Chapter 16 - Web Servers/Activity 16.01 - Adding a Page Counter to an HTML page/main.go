package main

import (
	"fmt"
	"net/http"
)

type PageWithCounter struct {
	counter int
	header  string
	content string
}

func (p *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.counter += 1
	hdr := fmt.Sprintf("<h1>%s</h1>", p.header)
	cnt := fmt.Sprintf("<h3>%s<h3>", p.content)
	views := fmt.Sprintf("Views: %d", p.counter)

	w.Write([]byte(hdr))
	w.Write([]byte(cnt))
	w.Write([]byte(views))
}

func main() {
	mainPage := PageWithCounter{counter: 0, header: "Hello World", content: "This is the main page"}
	firstChapter := PageWithCounter{counter: 0, header: "Chapter 1", content: "This is the first chapter"}
	secondChapter := PageWithCounter{counter: 0, header: "Chapter 2", content: "This is the second chapter"}

	http.Handle("/", &mainPage)
	http.Handle("/chapter1", &firstChapter)
	http.Handle("/chapter2", &secondChapter)
	http.ListenAndServe(":8085", nil)
}
