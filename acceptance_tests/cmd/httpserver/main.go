package main

import (
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
