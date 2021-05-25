package handlers

import (
	"context"

	"github.com/google/uuid"

	pb "account"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AccountServer {
	return accountService{}
}

type accountService struct{}

func (s accountService) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	var resp pb.StatusResponse
	resp.Status = pb.ServiceStatus_OK
	return &resp, nil
}

func (s accountService) AuthLogin(ctx context.Context, in *pb.AuthLoginRequest) (*pb.AuthLoginResponse, error) {
	var resp pb.AuthLoginResponse
	resp.Token = uuid.NewString()
	return &resp, nil
}
