package parser

import (
	"encoding/json"
)

var (
	KeyPath     = "path"
	KeyTemplate = "template"
	KeyData     = "data"
	KeyRoutes   = "routes"
)

type Route struct {
	Path     string                 `json:"path"`
	Template string                 `json:"template"`
	Data     map[string]interface{} `json:"data"`
	Routes   []*Route               `json:"routes"`
}

func NewRoute(input map[string]interface{}) *Route {
	path := ""
	template := ""
	var data = make(map[string]interface{})
	var routes = make([]*Route, 0)

	for key, value := range input {
		switch key {
		case KeyPath:
			path = value.(string)
		case KeyTemplate:
			template = value.(string)
		case KeyData:
			data = value.(map[string]interface{})
		case KeyRoutes:
			items := value.([]interface{})
			for _, item := range items {
				r := NewRoute(item.(map[string]interface{}))
				routes = append(routes, r)
			}
		default:
			return nil
		}
	}

	return &Route{
		Path:     path,
		Template: template,
		Data:     data,
		Routes:   routes,
	}
}

func (r *Route) String() string {
	ret, _ := json.Marshal(r)
	return string(ret)
}
