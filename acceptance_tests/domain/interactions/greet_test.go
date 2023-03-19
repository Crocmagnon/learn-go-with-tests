package interactions_test

import (
	go_specs_greet "github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/domain/interactions"
	"github.com/Crocmagnon/learn-go-with-tests/acceptance_tests/specifications"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("specification", func(t *testing.T) {
		specifications.GreetSpecification(
			t,
			specifications.GreetAdapter(go_specs_greet.Greet),
		)
	})
	t.Run("default to world", func(t *testing.T) {
		got := go_specs_greet.Greet("")
		want := "Hello, world"

		assert.Equal(t, got, want)
	})
}

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(go_specs_greet.Curse),
	)
}
