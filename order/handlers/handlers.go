package handlers

import (
	"context"

	"github.com/google/uuid"

	pb "github.com/eedkevin/go-kit-boilerplate/order"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.OrderServer {
	return orderService{}
}

type orderService struct{}

func (s orderService) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	var resp pb.StatusResponse
	return &resp, nil
}

func (s orderService) PlaceOrder(ctx context.Context, in *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	var resp pb.PlaceOrderResponse
	resp.OrderId = uuid.NewString()
	return &resp, nil
}
