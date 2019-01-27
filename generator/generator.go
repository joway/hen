package generator

import (
	"github.com/joway/hen/parser"
	"io/ioutil"
)

type Generator struct {
	Config *parser.Config
}

func New(configFormat, configPath string) *Generator {
	p := parser.New(configFormat)
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}
	c, err := p.Parse(string(buf))
	if err != nil {
		return nil
	}

	return &Generator{
		Config: c,
	}
}

func (g *Generator) Generate(input string) {

}
