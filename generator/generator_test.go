package generator

import "testing"

func TestGenerator_Generate(t *testing.T) {
	g := New("yaml", `
routes:
  - path: '/'
    template: 'file://../example/post/index.html'
    data:
      title: "Home"
    routes:
      - path: '/post'
        template: 'file://../example/post/detail.html'
        data:
          type: 'post'
          arthur: 'Joway'
        routes:
          - path: '/post/1'
            data:
              title: 'This is title!'
              content: 'Hi | 中文'
`)
	g.Generate()
}
