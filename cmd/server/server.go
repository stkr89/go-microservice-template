package server

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/shopr-org/grpc-service-template/common"
	"github.com/shopr-org/grpc-service-template/config"
	"github.com/shopr-org/grpc-service-template/endpoints"
	"github.com/shopr-org/grpc-service-template/pb"
	"github.com/shopr-org/grpc-service-template/service"
	transport "github.com/shopr-org/grpc-service-template/transports"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func InitServer() {
	logger := common.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Log("message", ".env file not found", "err", err)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	StartServer(logger)

	level.Error(logger).Log("exit", <-errs)
}

func StartServer(logger log.Logger) {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	endpoint := endpoints.MakeEndpoints(service.NewMathServiceImpl())
	grpcServer := transport.NewGRPCServer(endpoint, logger)

	err = config.InitialDBMigration(config.NewDB())
	if err != nil {
		panic(err)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterMathServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("message", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()
}
