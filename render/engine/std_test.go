package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSTDEngine_Render(t *testing.T) {
	e := &STDEngine{}
	input := `
{{add .int 1}}
{{.str}}
{{ range $key, $value := .array }}{{ $key }} {{ $value }} {{ end }}
{{.array}}
{{.obj}}
{{max .obj.int 1000}}
{{.obj.str}}
`
	data := make(map[string]interface{})
	data["int"] = 10
	data["str"] = "test"
	data["array"] = [2]string{"1", "2"}
	data["obj"] = map[string]interface{}{"int": 1, "str": "test"}
	output, err := e.Render(data, input)

	assert.NoError(t, err)
	assert.Equal(t, `
11
test
0 1 1 2 
[1 2]
map[int:1 str:test]
1000
test
`, output)
}
