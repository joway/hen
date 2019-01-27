package engine

type TemplateEngine interface {
    Render(string) string
}
