package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYAMLParser_Parse(t *testing.T) {
	p := &YAMLParser{}
	config, err := p.Parse(`
routes:
  - path: '/404'
    template: 'file://404.html'
  - path: '/'
    template: 'file://post/index.html'
    routes:
      - path: '/post/404'
        template: 'file://post/404.html'
      - path: '/post'
        template: 'file://post/detail.html'
        data:
          type: 'post'
          arthur: 'Joway'
        routes:
          - path: '/post/1'
            data:
              title: 'This is title!'
              content: 'file://post/content/1.md'
`)
	assert.NoError(t, err)

	routes := config.Routes
	assert.True(t, len(routes) == 2)
	route := routes[1]
	assert.Equal(t, "/", route.Path)
	assert.Equal(t, "file://post/index.html", route.Template)
	assert.Equal(t, 2, len(route.Routes))
	assert.Equal(t, "/post/404", route.Routes[0].Path)
	assert.Equal(t, 1, len(route.Routes[1].Routes))
	assert.Equal(t, "/post/1", route.Routes[1].Routes[0].Path)
	assert.Equal(t, "This is title!", route.Routes[1].Routes[0].Data["title"])
}
