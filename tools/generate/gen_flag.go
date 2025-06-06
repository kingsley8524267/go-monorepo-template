package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func genFlag(appName string) error {
	filename := filepath.Join("apps", appName, "flag.go")
	err := os.MkdirAll(filepath.Dir(filename), 0750)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	const tmpl = `package {{.PackageName}}

import (
	"flag"
	"fmt"
	"os"
)

var (
	fVersion bool
)

func init() {
	flag.BoolVar(&fVersion, "v", false, "show version")
	flag.Parse()

	flagHandler()
}

func flagHandler() {
	if fVersion {
		fmt.Printf("%s: %s\n", name, version)
		os.Exit(0)
	}
}
`

	t := template.Must(template.New("app").Parse(tmpl))
	return t.Execute(f, map[string]string{"PackageName": toSnakeCase(appName)})
}
