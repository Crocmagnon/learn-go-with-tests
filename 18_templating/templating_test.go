package templating_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"io"
	templating "learn-go-with-tests/18_templating"
	"testing"
)

func TestRender(t *testing.T) {
	post := templating.Post{
		Title:       "Hello, world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	renderer, err := templating.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("convert a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err = renderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	post := templating.Post{
		Title:       "Hello, world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	renderer, err := templating.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderer.Render(io.Discard, post)
	}
}
