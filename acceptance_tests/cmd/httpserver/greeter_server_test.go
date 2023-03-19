package main_test

import (
	"context"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/adapters/httpserver"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../../.",
			Dockerfile:    "./acceptance_tests/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	driver := httpserver.Driver{BaseURL: "http://127.0.0.1:8080", Client: &client}
	specifications.GreetSpecification(t, driver)
}
