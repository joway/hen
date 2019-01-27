package render

import (
	"github.com/joway/hen/render/engine"
)

type Render struct {
	engine engine.TemplateEngine
}

func New(engine engine.TemplateEngine) *Render {
	return &Render{
		engine: engine,
	}
}
func (r *Render) Output(data map[string]interface{}, input string) (string, error) {
	return r.engine.Render(data, input)
}
