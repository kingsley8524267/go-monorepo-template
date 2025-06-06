package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func genConfig(appName string) error {
	filename := filepath.Join("internal", "config", fmt.Sprintf("%s.go", appName))
	err := os.MkdirAll(filepath.Dir(filename), 0750)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	const tmpl = `package config

type {{.ConfigName}} struct {
	Common ` + "`mapstructure:\",squash\"`" + `
}
`

	t := template.Must(template.New("app").Parse(tmpl))
	return t.Execute(f, map[string]string{"ConfigName": toPascalCase(appName)})
}
