package httpserver

import (
	"fmt"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/domain/interactions"
	"net/http"
)

func GreetHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Greet(name))
}

func CurseHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Curse(name))
}
