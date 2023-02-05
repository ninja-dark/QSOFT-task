package server

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	srv    http.Server
	logger *zap.Logger
}

func NewServer(addr string, h http.Handler, logger *zap.Logger) *Server {
	return &Server{
		srv: http.Server{
			Addr:              addr,
			Handler:           h,
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			ReadHeaderTimeout: 30 * time.Second,
		},
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		if err := s.srv.Shutdown(stopCtx); err != nil {

			s.logger.Error("Server Shutdown Failed")
		}
	}()

	return s.srv.ListenAndServe()
}
