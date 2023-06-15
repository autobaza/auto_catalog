package main

import (
	"fmt"
	"github.com/autobaza/auto_catalog/endpoints"
	catalog "github.com/autobaza/auto_catalog/protos"
	"github.com/autobaza/auto_catalog/repository"
	"github.com/autobaza/auto_catalog/service"
	"github.com/autobaza/auto_catalog/transports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var repo = repository.NewRepository()

	srv := service.NewService(logger, repo)
	endpts := endpoints.MakeEndpoints(srv)
	grpcServer := transports.NewGRPCServer(endpts)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		catalog.RegisterAutoCatalogServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
