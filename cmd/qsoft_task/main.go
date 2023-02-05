package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ninja-dark/QSOFT-task/internal/config"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/api/handler"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/api/routergin"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/server"
	"github.com/ninja-dark/QSOFT-task/internal/logic"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logger.Sync() }()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.Read()
	if err != nil {
		logger.Sugar().Fatalf("Cannot load config, due to error: %s", err.Error())
	}

	l := logic.New()
	hs := handler.NewHandlers(l)
	rt := routergin.NewRouterGin(hs)
	logger.Sugar().Infof("Starting Gateway server on port:%s", cfg.Addr)
	srv := server.NewServer(cfg.Addr, rt, logger)

	if err := srv.Start(ctx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Sugar().Fatalf("Failed to start server: %s", err.Error())
		}
	}
}
