package main

import (
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", httpserver.GreetHandler)
	mux.HandleFunc("/curse", httpserver.CurseHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
