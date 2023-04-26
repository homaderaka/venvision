package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"venvision/internal/config"
	"venvision/internal/product"
	"venvision/internal/product/service"
	"venvision/pkg/logging"
	"venvision/pkg/metrics"
)

type App struct {
	cfg        *config.Config
	logger     *logging.Logger
	service    service.Service
	router     *gin.Engine
	httpServer *http.Server
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewApp(cfg *config.Config, logger *logging.Logger) (app App, err error) {
	logger.Println("router initializing")

	// router w/o logger and recovery middleware
	router := gin.New()

	// TODO: add grpc router

	// add recovery middleware
	// recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	v1 := router.Group("/api/v1/")

	logger.Println("heartbeat metric initializing")
	metricHandler := metrics.Handler{}
	metricHandler.Register(v1)

	prodService, err := service.NewProductService(logger)
	if err != nil {
		return
	}
	// TODO: add storage to the service

	proxyHandler := product.Handler{
		Logger:         *logger,
		ProductService: prodService,
	}

	proxyPath := v1.Group("/product")
	proxyHandler.Register(proxyPath)

	sharedCtx, cancel := context.WithCancel(context.Background())

	return App{
		cfg,
		logger,
		prodService,
		router,
		nil,
		sharedCtx,
		cancel,
	}, nil
}

func (a *App) Run() {
	a.startHTTP()
}

func (a *App) startHTTP() {
	a.logger.Info("start HTTP")

	var listener net.Listener

	if a.cfg.ListenConfig.Type == config.ListenTypeSock {
		appDir, err := filepath.Abs(os.Args[0])
		if err != nil {
			a.logger.Fatal(err)
		}
		socketPath := path.Join(appDir, a.cfg.ListenConfig.SocketFile)
		a.logger.Infof("socket path: %s", socketPath)

		a.logger.Info("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	} else {
		a.logger.Infof("bind application to host: %s and port: %s", a.cfg.ListenConfig.BindIP, a.cfg.ListenConfig.Port)
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.ListenConfig.BindIP, a.cfg.ListenConfig.Port))
		if err != nil {
			a.logger.Fatal(err)
		}
	}

	a.httpServer = &http.Server{
		Handler:      a.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")
		default:
			a.logger.Fatal(err)
		}
	}
	err := a.httpServer.Shutdown(a.ctx)
	if err != nil {
		a.logger.Fatal(err)
	}

}

func (a *App) Shutdown() {
	a.logger.Info("Initiating graceful shutdown...")
	a.cancel()
}
