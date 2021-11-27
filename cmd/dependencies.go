package main

import (
	"github.com/go-kit/log"
	"github.com/google/wire"
	"github.com/shopr-org/grpc-service-template/config"
	"github.com/shopr-org/grpc-service-template/service"
	"os"
)

func initLogger() log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

func initAddService() service.MathServiceImpl {
	wire.Build(wire.NewSet(
		initLogger,
		config.ProvideDB,
		service.ProviderMathDaoImpl,
		service.ProviderMathServiceImpl,
	))
	return service.MathServiceImpl{}
}
