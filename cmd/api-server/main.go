package main

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/francisco3ferraz/go-nuts/internal/config"
	"github.com/francisco3ferraz/go-nuts/internal/logger"
	"github.com/francisco3ferraz/go-nuts/internal/server"
	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	log := logger.New(cfg.Environment)
	defer func() {
		_ = log.Sync()
	}()

	apiServer := server.New(cfg, log)

	log.Info("starting api server", zap.String("addr", cfg.HTTPAddr))
	err := apiServer.Start(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("server stopped with error", zap.Error(err))
		panic(err)
	}

	log.Info("api server stopped")
}
