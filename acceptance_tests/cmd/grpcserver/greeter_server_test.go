package main_test

import (
	"fmt"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/grpcserver"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port           = "50051"
		dockerfilePath = "./acceptance_tests/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, dockerfilePath, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
}
