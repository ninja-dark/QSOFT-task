package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"syscall"

	"github.com/joho/godotenv"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/api/handler"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/api/routergin"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/server"
	"github.com/ninja-dark/QSOFT-task/internal/logic"
	"go.uber.org/zap"
)


func main(){
	logger, err := zap.NewProduction()
	if err != nil{
		log.Fatal(err)
	}
	defer func() { _ = logger.Sync() }()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	if err := godotenv.Load("../config/.env"); err != nil {
		logger.Sugar().Fatalf("Cannot load config, due to error: %s", err.Error())
	}

	l := logic.New()
	hs := handler.NewHandlers(l)
	rt := routergin.NewRouterGin(hs)
	logger.Sugar().Infof("Starting Gateway server on port:%s", os.Getenv("GATE_PORT"))
	srv := server.NewServer(os.Getenv("GATE_PORT"), rt)

	srv.Start()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			logger.Info("cancel context")
			srv.Stop()
			cancel() //Если пришёл сигнал SigInt - завершаем контекст
			
		}
	}


}