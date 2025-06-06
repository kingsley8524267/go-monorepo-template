package my_app

import (
	"go-monorepo-template/internal/config"
	"go-monorepo-template/internal/logger"
	"go-monorepo-template/internal/signal"
	"golang.org/x/net/context"
)

const name = "my-app"

var version = "DEV"

type App struct {
	info   config.App
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.MyApp
}

func New() (*App, error) {
	info := config.App{
		Name:    name,
		Version: version,
	}

	var err error
	cfg := new(config.MyApp)
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
