package main

import (
	go_specs_greet "github.com/Crocmagnon/learn-go-with-tests/acceptance_tests"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
