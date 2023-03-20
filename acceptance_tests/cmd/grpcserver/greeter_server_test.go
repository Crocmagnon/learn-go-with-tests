package main_test

import (
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/grpcserver"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()

	var (
		port           = "50051"
		dockerfilePath = "./acceptance_tests/Dockerfile"
		endpoint       = adapters.StartDockerServer(t, dockerfilePath, port, "grpcserver", "")
		driver         = grpcserver.Driver{Addr: endpoint}
	)

	t.Cleanup(driver.Close)
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
