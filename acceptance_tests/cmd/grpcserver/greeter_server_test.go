package main_test

import (
	"fmt"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/grpcserver"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerfilePath = "./acceptance_tests/cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, dockerfilePath, port)
	specifications.GreetSpecification(t, driver)
}
