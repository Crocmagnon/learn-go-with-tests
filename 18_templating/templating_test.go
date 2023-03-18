package templating_test

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
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

	t.Run("convert a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := templating.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
