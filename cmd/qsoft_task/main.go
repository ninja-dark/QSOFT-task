package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"syscall"

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
	l := logic.New()
	hs := handler.NewHandlers(l)
	rt := routergin.NewRouterGin(hs)
	srv := server.NewServer(":8080", rt)

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