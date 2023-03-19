package main_test

import (
	"fmt"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/httpserver"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerfilePath = "./acceptance_tests/cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://127.0.0.1:%s", port)
		client         = http.Client{Timeout: 1 * time.Second}
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &client}
	)

	adapters.StartDockerServer(t, dockerfilePath, port)
	specifications.GreetSpecification(t, driver)
}
