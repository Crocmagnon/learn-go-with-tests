package templating_test

import (
	"bytes"
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
		err := templating.Render(&buf, post)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello, world</h1>
<p>This is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
