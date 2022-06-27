package test

import (
	"github.com/joho/godotenv"
	"github.com/stkr89/go-microservice-template/cmd/server"
	"github.com/stkr89/go-microservice-template/common"
	"github.com/stkr89/go-microservice-template/endpoints"
	"github.com/stkr89/go-microservice-template/service"
	transport "github.com/stkr89/go-microservice-template/transports"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type HTTPIntegrationTestSuite struct {
	suite.Suite
	handler http.Handler
}

func (suite *HTTPIntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../.env")
	suite.NoError(err)

	e := endpoints.MakeEndpoints(service.NewMathServiceImpl())
	server.StartServer(common.NewLogger(), e, false, true)

	suite.handler = transport.NewHTTPHandler(e)
}

// new test cases

func TestHTTPIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPIntegrationTestSuite))
}
