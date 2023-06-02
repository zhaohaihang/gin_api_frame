package http

import (
	"context"
	"fmt"
	"gin_api_frame/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.Server
}

func NewHttpServer(c *config.Config, r *gin.Engine) *HttpServer {
	return &HttpServer{
		config: c,
		router: r,
	}
}

var HttpServerProviderSet = wire.NewSet(NewHttpServer)

func (s *HttpServer) Start() error {
	s.httpServer = &http.Server{Addr: fmt.Sprintf("%s%s", s.config.Server.ServerHost, s.config.Server.ServerPort), Handler: s.router}
	logrus.Info("http server starting...")
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("start http server error %s", err.Error())
			return
		}
	}()
	return nil
}

func (s *HttpServer) Stop() error {
	logrus.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) 
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}
