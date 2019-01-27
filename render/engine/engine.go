package engine

type TemplateEngine interface {
	Render(data map[string]interface{}, input string) (string, error)
}
