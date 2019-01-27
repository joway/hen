package engine

import (
	"github.com/Masterminds/sprig"

	"bytes"
	"html/template"
)

type STDEngine struct {
	TemplateEngine
}

func (e *STDEngine) Render(data map[string]interface{}, input string) (string, error) {
	tmpl, err := template.New("base").Funcs(sprig.FuncMap()).Parse(input)
	if err != nil {
		return "", err
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	if err != nil {
		return "", err
	}
	return doc.String(), nil
}
