package parser

type Parser interface {
	Parse(string) (*Config, error)
}

func New(format string) Parser {
	switch format {
	case "toml":
		return &TOMLParser{}
	case "yaml":
		return &YAMLParser{}
	default:
		return nil
	}
}
