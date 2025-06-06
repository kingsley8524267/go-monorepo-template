package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func genApp(appName string) error {
	filename := filepath.Join("apps", appName, "app.go")
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
	"go-monorepo-template/internal/config"
	"go-monorepo-template/internal/logger"
	"go-monorepo-template/internal/signal"
	"golang.org/x/net/context"
)

const name = "{{.AppName}}"

var version = "DEV"

type App struct {
	info   config.App
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.{{.ConfigName}}
}

func New() (*App, error) {
	info := config.App{
		Name:    name,
		Version: version,
	}

	var err error
	cfg := new(config.{{.ConfigName}})
	if err = config.LoadConfig(info, cfg); err != nil {
		return nil, err
	}

	if err = logger.Init(cfg.Logger); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	signal.AddCloseFunc("logger", func() error { return logger.Sync() })
	return &App{
		info:   info,
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}, nil
}

func (s *App) Run() {
	logger.Infof("%s started", s.info.String())

	/*
		Add all your application logic here
	*/

	signal.WaitClose(s.info)
}

func (s *App) Close() error {
	s.cancel()
	return nil
}

func (s *App) String() string {
	return s.info.String()
}
`

	t := template.Must(template.New("app").Parse(tmpl))
	return t.Execute(f, map[string]string{"PackageName": toSnakeCase(appName), "AppName": appName, "ConfigName": toPascalCase(appName)})
}
