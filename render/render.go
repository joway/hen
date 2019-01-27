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

func (r *Render) Output(input string) string {
    return r.engine.Render(input)
}
