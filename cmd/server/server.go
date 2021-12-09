package server

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/shopr-org/grpc-service-template/common"
	"github.com/shopr-org/grpc-service-template/config"
	"github.com/shopr-org/grpc-service-template/endpoints"
	"github.com/shopr-org/grpc-service-template/pb"
	"github.com/shopr-org/grpc-service-template/service"
	transport "github.com/shopr-org/grpc-service-template/transports"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

	e := endpoints.MakeEndpoints(service.NewMathServiceImpl())
	StartServer(logger, e, true, true)

	level.Error(logger).Log("exit", <-errs)
}

func StartServer(logger log.Logger, e endpoints.Endpoints, startGRPC, startHTTP bool) {
	err := config.InitialDBMigration(config.NewDB())
	if err != nil {
		panic(err)
	}

	if startGRPC {
		startGRPCServer(logger, e)
	}

	if startHTTP {
		startHTTPServer(logger, e)
	}
}

func startHTTPServer(logger log.Logger, e endpoints.Endpoints) {
	listener, err := getListener(os.Getenv("PORT"))
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}

	httpHandler := transport.NewHTTPHandler(e)

	go func() {
		level.Info(logger).Log("msg", "Starting HTTP server 🚀")
		http.Serve(listener, httpHandler)
	}()
}

func startGRPCServer(logger log.Logger, endpoints endpoints.Endpoints) {
	listener, err := getListener(os.Getenv("GRPC_PORT"))
	if err != nil {
		logger.Log("transport", "GRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	grpcServer := transport.NewGRPCServer(endpoints)
	baseServer := grpc.NewServer()
	pb.RegisterMathServiceServer(baseServer, grpcServer)

	go func() {
		level.Info(logger).Log("msg", "Starting GRPC server 🚀")
		baseServer.Serve(listener)
	}()
}

func getListener(port string) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%s", port))
}
