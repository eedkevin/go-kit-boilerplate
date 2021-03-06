package handlers

import (
	"google.golang.org/grpc"

	accountPb "github.com/eedkevin/go-kit-boilerplate/account"
	accountGrpc "github.com/eedkevin/go-kit-boilerplate/account/svc/client/grpc"
)

var grpcClients *GRPCClients

type GRPCClients struct {
	Account accountPb.AccountServer
}

func init() {
	conn, err := grpc.Dial("localhost:5040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcClients = &GRPCClients{}
	grpcClients.Account, err = accountGrpc.New(conn)
	if err != nil {
		panic(err)
	}
}
