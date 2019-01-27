package generator

import (
	"encoding/json"
	"fmt"
	"github.com/joway/hen/parser"
	"github.com/joway/hen/render"
	"github.com/joway/hen/render/engine"
	"io/ioutil"
	"log"
	"strings"
)

type Generator struct {
	Config *parser.Config

	Render *render.Render
}

type File struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func (f *File) String() string {
	data, _ := json.Marshal(f)
	return string(data)
}

func New(configFormat string, configData string) *Generator {
	p := parser.New(configFormat)

	c, err := p.Parse(configData)
	if err != nil {
		return nil
	}

	templateEngine := engine.NewSTDEngine()

	return &Generator{
		Config: c,
		Render: render.New(templateEngine),
	}
}

func NewFromFile(configPath string) *Generator {
	configFormat := ""
	if strings.HasSuffix(configPath, ".yaml") || strings.HasSuffix(configPath, ".yml") {
		configFormat = "yaml"
	} else if strings.HasSuffix(configPath, ".toml") {
		configFormat = "toml"
	}
	if configFormat == "" {
		return nil
	}

	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}
	data := string(buf)

	return New(configFormat, data)
}

func cleanPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return path
}

func readProperty(property string) string {
	if strings.HasPrefix(property, "file://") {
		buf, err := ioutil.ReadFile(property[7:])
		if err != nil {
			log.Fatal(err)
		}
		return string(buf)
	}

	return property
}

func generate(render *render.Render, routes []*parser.Route) []File {
	files := make([]File, 0)
	for _, route := range routes {
		content := readProperty(route.Template)
		output, err := render.Output(route.Data, content)
		if err != nil {
			log.Fatal(err)
		}

		if route.Path != "" && route.Template != "" {
			path := cleanPath(route.Path)
			if path == "" {
				path = "index"
			}
			file := File{
				Path:    fmt.Sprintf("%s.html", path),
				Content: output,
			}
			files = append(files, file)
		}

		if len(route.Routes) > 0 {
			_files := generate(render, route.Routes)
			files = append(files, _files...)
		}
	}
	return files
}

func (g *Generator) Generate() {
	files := generate(g.Render, g.Config.Routes)
	fmt.Println(files)
}
