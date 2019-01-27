package parser

import "encoding/json"

type Config struct {
	Routes []*Route `json:"routes"`
}

func NewConfig(data map[string]interface{}) *Config {
	routes := NewRoute(data)

	return &Config{
		Routes: routes.Routes,
	}
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}
