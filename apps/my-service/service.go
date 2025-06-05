package my_service

import (
	"go-monorepo-template/internal/config"
	"go-monorepo-template/internal/logger"
	"go-monorepo-template/internal/signal"
	"golang.org/x/net/context"
)

const name = "my-service"

var version = "DEV"

type Service struct {
	info   config.Service
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.MyService
}

func New() (*Service, error) {
	info := config.Service{
		Name:    name,
		Version: version,
	}

	var err error
	cfg := new(config.MyService)
	if err = config.LoadConfig(info, cfg); err != nil {
		return nil, err
	}

	if err = logger.Init(cfg.Logger); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	signal.AddCloseFunc("logger", func() error { return logger.Sync() })
	return &Service{
		info:   info,
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}, nil
}

func (s *Service) Run() {
	logger.Infof("%s started", s.info.String())

	/*
		Add all your application logic here
	*/

	signal.WaitClose(s.info)
}

func (s *Service) Close() error {
	s.cancel()
	return nil
}

func (s *Service) String() string {
	return s.info.String()
}
