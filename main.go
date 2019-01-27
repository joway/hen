package main

import (
	"fmt"
	"github.com/joway/hen/generator"
)

func main() {
	configFormat := "toml"
	configPath := "example/config.toml"
	hen := generator.New(configFormat, configPath)
	fmt.Println(hen.Config)

    configFormat = "yaml"
    configPath = "example/config.yml"
    hen = generator.New(configFormat, configPath)
    fmt.Println(hen.Config)
}
