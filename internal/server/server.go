package server

import (
	"context"
	"net/http"
	"time"

	"github.com/francisco3ferraz/go-nuts/internal/api/router"
	"github.com/francisco3ferraz/go-nuts/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
	log             *zap.Logger
}

func New(cfg config.Config, log *zap.Logger) *Server {
	apiHandler := router.New()

	httpServer := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           apiHandler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	return &Server{
		httpServer:      httpServer,
		shutdownTimeout: cfg.ShutdownTimeout,
		log:             log,
	}
}

func (s *Server) Start(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.httpServer.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		s.log.Info("shutting down server")
		return s.httpServer.Shutdown(timeoutCtx)
	case err := <-errCh:
		return err
	}
}
