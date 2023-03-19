package go_specs_greet_test

import (
	go_specs_greet "github.com/Crocmagnon/learn-go-with-tests/acceptance_tests"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_specs_greet.Greet),
	)
}
