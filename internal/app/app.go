package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Flowwzen/api-gateway-service/internal/config"
	"github.com/Flowwzen/api-gateway-service/internal/router"
	"github.com/Flowwzen/api-gateway-service/internal/server"
)

func Run() {
	cfg := config.Load()

	handler := router.New()

	srv := server.New(
		handler,
		cfg.Server.Port,
		cfg.Server.ReadTimeout,
		cfg.Server.WriteTimeout,
	)

	go srv.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

  	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = srv.Shutdown(ctx)
}
