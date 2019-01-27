package parser

import (
	"github.com/ghodss/yaml"
)

type YAMLParser struct {
	Parser
}

func (p *YAMLParser) Parse(input string) (*Config, error) {
	tree := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(input), &tree)
	if err != nil {
		return nil, err
	}

	return NewConfig(tree), nil
}
