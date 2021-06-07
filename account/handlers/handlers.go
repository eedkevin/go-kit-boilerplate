package handlers

import (
	"context"

	"github.com/google/uuid"

	pb "github.com/eedkevin/go-kit-boilerplate/account"
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

func (s accountService) AuthToken(ctx context.Context, in *pb.AuthTokenRequest) (*pb.AuthTokenResponse, error) {
	var resp pb.AuthTokenResponse
	resp.Token = uuid.NewString()
	return &resp, nil
}

func (s accountService) AuthTokenValidate(ctx context.Context, in *pb.AuthTokenValidateRequest) (*pb.AuthTokenValidateResponse, error) {
	var resp pb.AuthTokenValidateResponse
	resp.IsValid = true
	return &resp, nil
}
