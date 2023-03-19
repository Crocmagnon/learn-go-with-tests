package specifications

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Gab")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Gab")
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t testing.TB, meany MeanGreeter) {
	got, err := meany.Curse("Gab")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, Gab!")
}
