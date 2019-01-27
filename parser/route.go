package parser

import (
	"encoding/json"
	"github.com/imdario/mergo"
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
	path, existed := input[KeyPath].(string)
	if !existed {
		path = "/"
	}
	template, existed := input[KeyTemplate].(string)
	if !existed {
		template = ""
	}
	data, existed := input[KeyData].(map[string]interface{})
	if !existed {
		data = make(map[string]interface{})
	}
	routes := make([]*Route, 0)
	_routes, existed := input[KeyRoutes].([]interface{})
	if existed {
		for _, _route := range _routes {
			subRoutes := _route.(map[string]interface{})
			// inherit properties
			if _, existed := subRoutes[KeyPath]; !existed {
				subRoutes[KeyPath] = path
			}
			if _, existed := subRoutes[KeyTemplate]; !existed {
				subRoutes[KeyTemplate] = template
			}
			if _, existed := subRoutes[KeyData]; !existed {
				subRoutes[KeyData] = data
			} else {
				dst := subRoutes[KeyData].(map[string]interface{})
				if err := mergo.Merge(&dst, data); err != nil {
					return nil
				}
				subRoutes[KeyData] = dst
			}
			r := NewRoute(subRoutes)
			routes = append(routes, r)
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
