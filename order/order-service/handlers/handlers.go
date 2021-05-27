package handlers

import (
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	geoPb "geo"
	geoGrpc "geo/geo-service/svc/client/grpc"
	pb "order"
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

	conn, err := grpc.Dial("geo-service:5040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client, err := geoGrpc.New(conn)
	if err != nil {
		panic(err)
	}

	r, err := client.Distance(ctx, &geoPb.DistanceRequest{
		Origin:      &geoPb.GeoCoordinate{Lat: 123.1, Lng: 213.2},
		Destination: &geoPb.GeoCoordinate{Lat: 123.1, Lng: 213.2},
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", r)

	resp.OrderId = uuid.NewString()
	return &resp, nil
}
