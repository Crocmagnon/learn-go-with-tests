package interactions_test

import (
	go_specs_greet "github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/domain/interactions"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_specs_greet.Greet),
	)
}

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(go_specs_greet.Curse),
	)
}
