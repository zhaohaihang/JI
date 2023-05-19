package server

import (
	"context"
	"fmt"
	"ji/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

)

type Server struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(c *config.Config, r *gin.Engine) *Server {
	return &Server{
		config: c,
		router: r,
	}
}

var ServerProviderSet = wire.NewSet(NewServer)

func (s *Server) Start() error {
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

func (s *Server) Stop() error {
	logrus.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 平滑关闭,等待5秒钟处理
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

func (s *Server) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	signal := <-c
	logrus.Infof("receive a signal signal:%s", signal.String())
	if s.httpServer != nil {
		if err := s.Stop(); err != nil {
			logrus.Warnf("stop http server error:%s", err.Error())
		}
	}

	os.Exit(0)
}
