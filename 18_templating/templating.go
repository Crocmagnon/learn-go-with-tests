package templating

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
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
	if err := r.templ.Execute(w, p); err != nil {
		return err
	}
	return nil
}
