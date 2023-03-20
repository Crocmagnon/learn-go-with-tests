package main_test

import (
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/httpserver"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	var (
		port           = "8080"
		dockerfilePath = "./acceptance_tests/Dockerfile"
		client         = http.Client{Timeout: 1 * time.Second}
		endpoint       = adapters.StartDockerServer(t, dockerfilePath, port, "httpserver", "http")
		driver         = httpserver.Driver{BaseURL: endpoint, Client: &client}
	)

	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
