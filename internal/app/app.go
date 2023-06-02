package app

import (
	"gin_api_frame/config"
	"gin_api_frame/internal/cron"
	"gin_api_frame/internal/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type App struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.HttpServer
	cronServer *cron.CronServer
}

func NewApp(c *config.Config, r *gin.Engine, hs *http.HttpServer, cs *cron.CronServer) *App {
	return &App{
		config:     c,
		router:     r,
		httpServer: hs,
		cronServer: cs,
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
	os.Exit(0)
}
