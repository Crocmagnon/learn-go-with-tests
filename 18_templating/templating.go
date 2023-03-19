package templating

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
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

func (p Post) RenderBody() template.HTML {
	md := []byte(p.Body)
	md = markdown.NormalizeNewlines(md)

	extensions := parser.CommonExtensions
	markdownParser := parser.NewWithExtensions(extensions)

	doc := markdownParser.Parse(md)

	flags := html.CommonFlags
	opts := html.RendererOptions{Flags: flags}
	renderer := html.NewRenderer(opts)

	rendered := markdown.Render(doc, renderer)
	return template.HTML(string(rendered))
}

func (p Post) SlugTitle() string {
	title := strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
	nonSlug := regexp.MustCompile("[^a-z-_]")
	return nonSlug.ReplaceAllString(title, "-")
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
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
