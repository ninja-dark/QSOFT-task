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

func NewServer(addr string, h http.Handler) *Server {
        s := &Server{}

        s.srv = http.Server{
                Addr:              addr,
                Handler:           h,
                ReadTimeout:       30 * time.Second,
                WriteTimeout:      30 * time.Second,
                ReadHeaderTimeout: 30 * time.Second,
        }

        return s
}

/*func (s *Server) Stop() {
        ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
        if err := s.srv.Shutdown(ctx); err != nil {
                s.logger.Info("Server Shutdown Failed")
        }
        cancel()
}*/

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