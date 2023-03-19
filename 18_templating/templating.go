package templating

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"regexp"
	"strings"
)

var (
	//go:embed "templates"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newRenderedPost(p))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", newRenderedPosts(posts))
}

type renderedPost struct {
	Body      template.HTML
	SlugTitle string
	RawPost   Post
}

func newRenderedPost(post Post) renderedPost {
	md := []byte(post.Body)
	md = markdown.NormalizeNewlines(md)

	extensions := parser.CommonExtensions
	// The parser must not be reused, or it can cause a panic
	// see https://github.com/gomarkdown/markdown/issues/229
	p := parser.NewWithExtensions(extensions)

	html := markdown.ToHTML(md, p, nil)
	body := template.HTML(html)

	title := strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
	nonSlug := regexp.MustCompile("[^a-z-_]")
	title = nonSlug.ReplaceAllString(title, "-")

	return renderedPost{Body: body, RawPost: post, SlugTitle: title}
}

func newRenderedPosts(posts []Post) []renderedPost {
	var rendered []renderedPost
	for _, post := range posts {
		r := newRenderedPost(post)
		rendered = append(rendered, r)
	}
	return rendered
}
