package app

import (
	"ji/config"
	"ji/internal/backproc"
	"ji/internal/cron"
	"ji/internal/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type App struct {
	config         *config.Config
	router         *gin.Engine
	httpServer     *http.HttpServer
	cronServer     *cron.CronServer
	backProcServer *backproc.BackProcServer
	logger       *logrus.Logger
}

func NewApp(c *config.Config,
	r *gin.Engine,
	hs *http.HttpServer,
	cs *cron.CronServer,
	bps *backproc.BackProcServer,
	l *logrus.Logger) *App {
	return &App{
		config:         c,
		router:         r,
		httpServer:     hs,
		cronServer:     cs,
		backProcServer: bps,
		logger:       l,
	}
}

var AppProviderSet = wire.NewSet(NewApp)

func (a *App) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	if a.cronServer != nil {
		a.cronServer.Start()
	}

	if a.backProcServer != nil {
		err := a.backProcServer.Start()
		return errors.Wrap(err, "backProc Server start error")
	}

	return nil
}

func (a *App) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	logrus.Infof("receive a signal: %s", s.String())
	if a.httpServer != nil {
		if err := a.httpServer.Stop(); err != nil {
			logrus.Warn("stop http server error %s", err)
		}
	}

	if a.cronServer != nil {
		a.cronServer.Stop()
	}
	
	if a.backProcServer != nil {
		a.backProcServer.Stop()
	}
	os.Exit(0)
}
