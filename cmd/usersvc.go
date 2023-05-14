package main

import (
	"context"
	"flag"
	"github.com/morris-zheng/go-slim-micro-usersvc/internal/conf"
	"github.com/morris-zheng/go-slim-micro-usersvc/internal/delivery"
	"github.com/morris-zheng/go-slim-micro-usersvc/internal/domain"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configPath := flag.String("f", "config.yaml", "config file path")
	flag.Parse()
	c := conf.Load(*configPath)

	svc := domain.NewServiceContext(c)

	server := delivery.NewGrpcServer()
	server.Register(svc)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-quit
		cancel()
	}()

	server.Run(ctx, svc)

	wait, waitCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer waitCancel()
	<-wait.Done()
}
