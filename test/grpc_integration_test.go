package test

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stkr89/go-microservice-template/cmd/server"
	"github.com/stkr89/go-microservice-template/common"
	"github.com/stkr89/go-microservice-template/endpoints"
	"github.com/stkr89/go-microservice-template/pb"
	"github.com/stkr89/go-microservice-template/service"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"os"
	"testing"
)

type GRPCIntegrationTestSuite struct {
	suite.Suite
	conn   *grpc.ClientConn
	client pb.MathServiceClient
}

func (suite *GRPCIntegrationTestSuite) SetupSuite() {
	err := godotenv.Load("../.env")
	suite.NoError(err)

	e := endpoints.MakeEndpoints(service.NewMathServiceImpl())
	server.StartServer(common.NewLogger(), e, true, false)

	conn, err := grpc.Dial(fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")), grpc.WithInsecure())
	suite.NoError(err)
	suite.conn = conn
	suite.client = pb.NewMathServiceClient(conn)
}

func (suite *GRPCIntegrationTestSuite) TearDownTestSuite() {
	_ = suite.conn.Close()
}

// new test cases

func TestGRPCIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(GRPCIntegrationTestSuite))
}
