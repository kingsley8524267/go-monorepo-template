package signal

import (
	"go-monorepo-template/internal/config"
	"go-monorepo-template/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

type closeFunc struct {
	name  string
	close func() error
}

var cfs []closeFunc

func AddCloseFunc(name string, f func() error) {
	cfs = append(cfs, closeFunc{name: name, close: f})
}

func WaitClose(service config.Service) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infof("%s stopping...", service.Name)
	for _, cf := range cfs {
		logger.Infof("%s executing close function: %s", service.Name, cf.name)
		if err := cf.close(); err != nil {
			logger.Warnf("%s close function %s error: %v", service.Name, cf.name, err)
		}
	}
	logger.Infof("%s stopped!", service.Name)
}
