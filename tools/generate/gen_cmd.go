package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func genCmd(appName string) error {
	filename := filepath.Join("cmd", appName, "main.go")
	err := os.MkdirAll(filepath.Dir(filename), 0750)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	const tmpl = `package main

import "go-monorepo-template/apps/{{.AppName}}"

func main() {
	app, err := {{.PackageName}}.New()
	if err != nil {
		panic(err)
	}

	app.Run()
}
`

	t := template.Must(template.New("app").Parse(tmpl))
	return t.Execute(f, map[string]string{"AppName": appName, "PackageName": toSnakeCase(appName)})
}
