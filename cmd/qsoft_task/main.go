package main

import (
        "context"
        "errors"
        "log"
        "net/http"
        "os"
        "os/signal"

        "github.com/joho/godotenv"
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

        if err := godotenv.Load(".env"); err != nil {
                logger.Sugar().Fatalf("Cannot load config, due to error: %s", err.Error())
        }

        l := logic.New()
        hs := handler.NewHandlers(l)
        rt := routergin.NewRouterGin(hs)
        logger.Sugar().Infof("Starting Gateway server on port:%s", os.Getenv("GATE_PORT"))
        srv := server.NewServer(os.Getenv("GATE_PORT"), rt)

        if err := srv.Start(ctx); err != nil {
                if !errors.Is(err, http.ErrServerClosed) {
                        logger.Sugar().Fatalf("Failed to start server: %s", err.Error())
                }
        }
}