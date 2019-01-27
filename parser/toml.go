package parser

import "github.com/pelletier/go-toml"

type TOMLParser struct {
	Parser
}

func (p *TOMLParser) Parse(input string) (*Config, error) {
	tree, err := toml.Load(input)
	if err != nil {
		return nil, err
	}

	return NewConfig(tree.ToMap()), nil
}
